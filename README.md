# Prometheus Hetzner SD

[![Build Status](http://cloud.drone.io/api/badges/promhippie/prometheus-hetzner-sd/status.svg)](http://cloud.drone.io/promhippie/prometheus-hetzner-sd)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/87cbb93f28be43a2a871018f106bc286)](https://www.codacy.com/app/promhippie/prometheus-hetzner-sd?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=promhippie/prometheus-hetzner-sd&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/promhippie/prometheus-hetzner-sd?status.svg)](http://godoc.org/github.com/promhippie/prometheus-hetzner-sd)
[![Go Report](http://goreportcard.com/badge/github.com/promhippie/prometheus-hetzner-sd)](http://goreportcard.com/report/github.com/promhippie/prometheus-hetzner-sd)
[![](https://images.microbadger.com/badges/image/promhippie/prometheus-hetzner-sd.svg)](http://microbadger.com/images/promhippie/prometheus-hetzner-sd "Get your own image badge on microbadger.com")

This project provides a server to automatically discover nodes within your Hetzner account in a Prometheus SD compatible format.

## Install

You can download prebuilt binaries from our [GitHub releases](https://github.com/promhippie/prometheus-hetzner-sd/releases), or you can use our Docker images published on [Docker Hub](https://hub.docker.com/r/promhippie/prometheus-hetzner-sd/tags/). If you need further guidance how to install this take a look at our [documentation](https://promhippie.github.io/prometheus-hetzner-sd/#getting-started).

## Integration

### Available labels

The following list of meta labels can be used to relabel your scrape results entirely. Hopefully the names are self-explaining, that's why I have skipped a description for each label.

* `__address__`
* `__meta_hetzner_cancelled`
* `__meta_hetzner_dc`
* `__meta_hetzner_flatrate`
* `__meta_hetzner_ipv4`
* `__meta_hetzner_name`
* `__meta_hetzner_number`
* `__meta_hetzner_product`
* `__meta_hetzner_project`
* `__meta_hetzner_status`
* `__meta_hetzner_throttled`
* `__meta_hetzner_traffic`

### Prometheus config

Here you get a snippet for the Prometheus `scrape_config` that configures Prometheus to scrape `node_exporter` assuming that it is deployed on all your servers.

```
- job_name: node
  file_sd_configs:
  - files: [ "/etc/prometheus/hetzner.json" ]
  relabel_configs:
  - source_labels: [__meta_hetzner_ipv4]
    replacement: "${1}:9100"
    target_label: __address__
  - source_labels: [__meta_hetzner_dc]
    target_label: datacenter
  - source_labels: [__meta_hetzner_name]
    target_label: instance
```

## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.11.

```bash
git clone https://github.com/promhippie/prometheus-hetzner-sd.git
cd prometheus-hetzner-sd

make generate build

./bin/prometheus-hetzner-sd -h
```

## Security

If you find a security issue please contact thomas@webhippie.de first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

* [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
