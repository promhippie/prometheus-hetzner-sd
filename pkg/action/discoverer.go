package action

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/appscode/go-hetzner"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery/targetgroup"
)

var (
	// providerPrefix defines the general prefix for all labels.
	providerPrefix = model.MetaLabelPrefix + "hetzner_"

	// Labels defines all available labels for this provider.
	Labels = map[string]string{
		"cancelled": providerPrefix + "cancelled",
		"dc":        providerPrefix + "dc",
		"flatrate":  providerPrefix + "flatrate",
		"ip":        providerPrefix + "ipv4",
		"name":      providerPrefix + "name",
		"number":    providerPrefix + "number",
		"product":   providerPrefix + "product",
		"project":   providerPrefix + "project",
		"status":    providerPrefix + "status",
		"throttled": providerPrefix + "throttled",
		"traffic":   providerPrefix + "traffic",
	}
)

// Discoverer implements the Prometheus discoverer interface.
type Discoverer struct {
	clients map[string]*hetzner.Client
	logger  log.Logger
	refresh int
	lasts   map[string]struct{}
}

// Run initializes fetching the targets for service discovery.
func (d Discoverer) Run(ctx context.Context, ch chan<- []*targetgroup.Group) {
	ticker := time.NewTicker(time.Duration(d.refresh) * time.Second)

	for {
		targets, err := d.getTargets(ctx)

		if err == nil {
			ch <- targets
		}

		select {
		case <-ticker.C:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func (d *Discoverer) getTargets(_ context.Context) ([]*targetgroup.Group, error) {
	current := make(map[string]struct{})
	targets := make([]*targetgroup.Group, 0)

	for project, client := range d.clients {
		now := time.Now()
		servers, _, err := client.Server.ListServers()
		requestDuration.WithLabelValues(project).Observe(time.Since(now).Seconds())

		if err != nil {
			level.Warn(d.logger).Log(
				"msg", "Failed to fetch servers",
				"project", project,
				"err", err,
			)

			requestFailures.WithLabelValues(project).Inc()
			continue
		}

		level.Debug(d.logger).Log(
			"msg", "Requested servers",
			"project", project,
			"count", len(servers),
		)

		for _, server := range servers {
			target := &targetgroup.Group{
				Source: fmt.Sprintf("hetzner/%d", server.ServerNumber),
				Targets: []model.LabelSet{
					{
						model.AddressLabel: model.LabelValue(server.ServerIP),
					},
				},
				Labels: model.LabelSet{
					model.AddressLabel:                   model.LabelValue(server.ServerIP),
					model.LabelName(Labels["project"]):   model.LabelValue(project),
					model.LabelName(Labels["name"]):      model.LabelValue(server.ServerName),
					model.LabelName(Labels["number"]):    model.LabelValue(strconv.Itoa(int(server.ServerNumber))),
					model.LabelName(Labels["ip"]):        model.LabelValue(server.ServerIP),
					model.LabelName(Labels["product"]):   model.LabelValue(server.Product),
					model.LabelName(Labels["dc"]):        model.LabelValue(strings.ToLower(server.Dc)),
					model.LabelName(Labels["traffic"]):   model.LabelValue(server.Traffic),
					model.LabelName(Labels["flatrate"]):  model.LabelValue(strconv.FormatBool(server.Flatrate)),
					model.LabelName(Labels["status"]):    model.LabelValue(server.Status),
					model.LabelName(Labels["throttled"]): model.LabelValue(strconv.FormatBool(server.Throttled)),
					model.LabelName(Labels["cancelled"]): model.LabelValue(strconv.FormatBool(server.Cancelled)),
				},
			}

			level.Debug(d.logger).Log(
				"msg", "Server added",
				"project", project,
				"source", target.Source,
			)

			current[target.Source] = struct{}{}
			targets = append(targets, target)
		}

	}

	for k := range d.lasts {
		if _, ok := current[k]; !ok {
			level.Debug(d.logger).Log(
				"msg", "Server deleted",
				"source", k,
			)

			targets = append(
				targets,
				&targetgroup.Group{
					Source: k,
				},
			)
		}
	}

	d.lasts = current
	return targets, nil
}
