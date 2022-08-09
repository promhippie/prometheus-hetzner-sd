# Prometheus Hetzner SD

[![Current Tag](https://img.shields.io/github/v/tag/promhippie/prometheus-hetzner-sd?sort=semver)](https://github.com/promhippie/prometheus-scw-sd) [![General Build](https://github.com/promhippie/prometheus-hetzner-sd/actions/workflows/general.yml/badge.svg)](https://github.com/promhippie/prometheus-hetzner-sd/actions/workflows/general.yaml) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/87cbb93f28be43a2a871018f106bc286)](https://www.codacy.com/gh/promhippie/prometheus-hetzner-sd/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=promhippie/prometheus-hetzner-sd&amp;utm_campaign=Badge_Grade) [![Go Doc](https://godoc.org/github.com/promhippie/prometheus-hetzner-sd?status.svg)](http://godoc.org/github.com/promhippie/prometheus-hetzner-sd) [![Go Report](http://goreportcard.com/badge/github.com/promhippie/prometheus-hetzner-sd)](http://goreportcard.com/report/github.com/promhippie/prometheus-hetzner-sd)

This project provides a server to automatically discover nodes within your
Hetzner account in a Prometheus SD compatible format.

## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our containers published on [Docker Hub][dockerhub] and [Quay][quayio].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.17, at least that's the version we are using.

```console
git clone https://github.com/promhippie/prometheus-hetzner-sd.git
cd prometheus-hetzner-sd

make generate build

./bin/prometheus-hetzner-sd -h
```

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Security

If you find a security issue please contact thomas@webhippie.de first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[releases]: https://github.com/promhippie/prometheus-hetzner-sd/releases
[dockerhub]: https://hub.docker.com/r/promhippie/prometheus-hetzner-sd/tags/
[quayio]: https://quay.io/repository/promhippie/prometheus-hetzner-sd?tab=tags
[docs]: https://promhippie.github.io/prometheus-hetzner-sd/#getting-started
[golang]: http://golang.org/doc/install.html
