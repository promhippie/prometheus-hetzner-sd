---
title: "Usage"
date: 2022-07-21T00:00:00+00:00
anchor: "getting-started"
weight: 10
---

## Installation

We won't cover further details how to properly setup [Prometheus][prometheus]
itself, we will only cover some basic setup based on [docker-compose][compose].
But if you want to run this service discovery without [docker-compose][compose]
you should be able to adopt that to your needs.

First of all we need to prepare a configuration for [Prometheus][prometheus]
that includes the service discovery which simply maps to a node exporter.

{{< highlight yaml >}}
global:
  scrape_interval: 1m
  scrape_timeout: 10s
  evaluation_interval: 1m

scrape_configs:
- job_name: node
  file_sd_configs:
  - files: [ "/etc/sd/hetzner.json" ]
  relabel_configs:
  - source_labels: [__meta_hetzner_public_ipv4]
    replacement: "${1}:9100"
    target_label: __address__
  - source_labels: [__meta_hetzner_dc]
    target_label: datacenter
  - source_labels: [__meta_hetzner_name]
    target_label: instance
- job_name: hetzner-sd
  static_configs:
  - targets:
    - hetzner-sd:9000
{{< / highlight >}}

After preparing the configuration we need to create the `docker-compose.yml`
within the same folder, this `docker-compose.yml` starts a simple
[Prometheus][prometheus] instance together with the service discovery. Don't
forget to update the envrionment variables with the required credentials. If you
are using a different volume for the service discovery you have to make sure
that the container user is allowed to write to this volume.

{{< highlight yaml >}}
version: '2'

volumes:
  prometheus:

services:
  prometheus:
    image: prom/prometheus:latest
    restart: always
    ports:
      - 9090:9090
    volumes:
      - prometheus:/prometheus
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
      - ./service-discovery:/etc/sd

  hetzner-sd:
    image: promhippie/prometheus-hetzner-sd:latest
    restart: always
    environment:
      - PROMETHEUS_HETZNER_LOG_PRETTY=true
      - PROMETHEUS_HETZNER_OUTPUT_FILE=/etc/sd/hetzner.json
      - PROMETHEUS_HETZNER_USERNAME=octocat
      - PROMETHEUS_HETZNER_PASSWORD=p455w0rd
    volumes:
      - ./service-discovery:/etc/sd
{{< / highlight >}}

Since our `latest` tag always refers to the `master` branch of the Git
repository you should always use some fixed version. You can see all available
tags at [DockerHub][dockerhub] or [Quay][quayio], there you will see that we
also provide a manifest, you can easily start the exporter on various
architectures without any change to the image name. You should apply a change
like this to the `docker-compose.yml` file:

{{< highlight diff >}}
  hetzner-sd:
-   image: promhippie/prometheus-hetzner-sd:latest
+   image: promhippie/prometheus-hetzner-sd:0.5.0
    restart: always
    environment:
      - PROMETHEUS_HETZNER_LOG_PRETTY=true
      - PROMETHEUS_HETZNER_OUTPUT_FILE=/etc/sd/hetzner.json
      - PROMETHEUS_HETZNER_USERNAME=octocat
      - PROMETHEUS_HETZNER_PASSWORD=p455w0rd
    volumes:
      - ./service-discovery:/etc/sd
{{< / highlight >}}

Depending on how you have launched and configured [Prometheus][prometheus] it's
possible that it's running as user `nobody`, in that case you should run the
service discovery as this user as well, otherwise [Prometheus][prometheus] won't
be able to read the generated JSON file:

{{< highlight diff >}}
  hetzner-sd:
    image: promhippie/prometheus-hetzner-sd:latest
    restart: always
+   user: '65534'
    environment:
      - PROMETHEUS_HETZNER_LOG_PRETTY=true
      - PROMETHEUS_HETZNER_OUTPUT_FILE=/etc/sd/hetzner.json
      - PROMETHEUS_HETZNER_USERNAME=octocat
      - PROMETHEUS_HETZNER_PASSWORD=p455w0rd
    volumes:
      - ./service-discovery:/etc/sd
{{< / highlight >}}

If you want to secure the access to the exporter or also the HTTP service
discovery endpoint you can provide a web config. You just need to provide a path
to the config file in order to enable the support for it, for details about the
config format look at the [documentation](#web-configuration) section:

{{< highlight diff >}}
  hetzner-sd:
    image: promhippie/prometheus-hetzner-sd:latest
    restart: always
    environment:
+     - PROMETHEUS_HETZNER_WEB_CONFIG=path/to/web-config.json
      - PROMETHEUS_HETZNER_LOG_PRETTY=true
      - PROMETHEUS_HETZNER_OUTPUT_FILE=/etc/sd/hetzner.json
      - PROMETHEUS_HETZNER_USERNAME=octocat
      - PROMETHEUS_HETZNER_PASSWORD=p455w0rd
    volumes:
      - ./service-discovery:/etc/sd
{{< / highlight >}}

To avoid the dependency on a shared filesystem between this service discovery
and the [Prometheus][prometheus] configuration directory, you are able to use
the new [HTTP service discovery][httpsd] starting with
[Prometheus][prometheus] >= v2.28, you just need to switch the engine for this
service discovery:

{{< highlight diff >}}
  hetzner-sd:
    image: promhippie/prometheus-hetzner-sd:latest
    restart: always
    environment:
      - PROMETHEUS_HETZNER_LOG_PRETTY=true
+     - PROMETHEUS_HETZNER_OUTPUT_ENGINE=http
      - PROMETHEUS_HETZNER_OUTPUT_FILE=/etc/sd/hetzner.json
      - PROMETHEUS_HETZNER_USERNAME=octocat
      - PROMETHEUS_HETZNER_PASSWORD=p455w0rd
    volumes:
      - ./service-discovery:/etc/sd
{{< / highlight >}}

To use the HTTP service discovery you just need to change the
[Prometheus][prometheus] configuration mentioned above a little bit:

{{< highlight yaml >}}
scrape_configs:
- job_name: node
  http_sd_configs:
  - url: http://hetzner-sd:9000/sd
  relabel_configs:
  - source_labels: [__meta_hetzner_public_ipv4]
    replacement: "${1}:9100"
    target_label: __address__
  - source_labels: [__meta_hetzner_dc]
    target_label: datacenter
  - source_labels: [__meta_hetzner_name]
    target_label: instance
{{< / highlight >}}

Finally the service discovery should be configured fine, let's start this stack
with [docker-compose][compose]], you just need to execute `docker-compose up`
within the directory where you have stored `prometheus.yml` and
`docker-compose.yml`. That's all, the service discovery should be up and
running. You can access [Prometheus][prometheus] at
[http://localhost:9090](http://localhost:9090).

{{< figure src="service-discovery.png" title="Prometheus service discovery for Hetzner" >}}

## Configuration

### Envrionment variables

If you prefer to configure the service with environment variables you can see
the available variables below, in case you want to configure multiple accounts
with a single service you are forced to use the configuration file as the
environment variables are limited to a single account. As the service is pretty
lightweight you can even start an instance per account and configure it entirely
by the variables, it's up to you.

{{< partial "envvars.md" >}}

### Web Configuration

If you want to secure the service by TLS or by some basic authentication you can
provide a `YAML` configuration file whch follows the [Prometheus][prometheus]
toolkit format. You can see a full configration example within the
[toolkit documentation][toolkit].

### Configuration file

Especially if you want to configure multiple accounts within a single service
discovery you got to use the configuration file. So far we support the file
formats `JSON` and `YAML`, if you want to get a full example configuration just
take a look at [our repository][configs], there you can always see the latest
configuration format. These example configurations include all available
options, they also include the default values.

## Labels

{{< partial "labels.md" >}}

## Metrics

prometheus_hetzner_sd_request_duration_seconds{project}
: Histogram of latencies for requests to the Hetzner API

prometheus_hetzner_sd_request_failures_total{project}
: Total number of failed requests to the Hetzner API

[prometheus]: https://prometheus.io
[compose]: https://docs.docker.com/compose/
[dockerhub]: https://hub.docker.com/r/promhippie/prometheus-hetzner-sd/tags/
[quayio]: https://quay.io/repository/promhippie/prometheus-hetzner-sd?tab=tags
[httpsd]: https://prometheus.io/docs/prometheus/2.28/configuration/configuration/#http_sd_config
[toolkit]: https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-configuration.md
[configs]: https://github.com/promhippie/prometheus-hetzner-sd/tree/master/config
