package action

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/appscode/go-hetzner"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery/targetgroup"
)

const (
	hetznerPrefix  = model.MetaLabelPrefix + "hetzner_"
	projectLabel   = hetznerPrefix + "project"
	nameLabel      = hetznerPrefix + "name"
	numberLabel    = hetznerPrefix + "number"
	ipLabel        = hetznerPrefix + "ipv4"
	productLabel   = hetznerPrefix + "product"
	dcLabel        = hetznerPrefix + "dc"
	trafficLabel   = hetznerPrefix + "traffic"
	flatrateLabel  = hetznerPrefix + "flatrate"
	statusLabel    = hetznerPrefix + "status"
	throttledLabel = hetznerPrefix + "throttled"
	cancelledLabel = hetznerPrefix + "cancelled"
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

func (d *Discoverer) getTargets(ctx context.Context) ([]*targetgroup.Group, error) {
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
			return nil, err
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
					model.AddressLabel:              model.LabelValue(server.ServerIP),
					model.LabelName(projectLabel):   model.LabelValue(project),
					model.LabelName(nameLabel):      model.LabelValue(server.ServerName),
					model.LabelName(numberLabel):    model.LabelValue(strconv.Itoa(int(server.ServerNumber))),
					model.LabelName(ipLabel):        model.LabelValue(server.ServerIP),
					model.LabelName(productLabel):   model.LabelValue(server.Product),
					model.LabelName(dcLabel):        model.LabelValue(strings.ToLower(server.Dc)),
					model.LabelName(trafficLabel):   model.LabelValue(server.Traffic),
					model.LabelName(flatrateLabel):  model.LabelValue(strconv.FormatBool(server.Flatrate)),
					model.LabelName(statusLabel):    model.LabelValue(server.Status),
					model.LabelName(throttledLabel): model.LabelValue(strconv.FormatBool(server.Throttled)),
					model.LabelName(cancelledLabel): model.LabelValue(strconv.FormatBool(server.Cancelled)),
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
