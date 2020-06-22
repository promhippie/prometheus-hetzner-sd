Bugfix: Binaries are not static linked

We fixed building properly static linked binaries, since the last release and a
major refactoring of the binaries and the CI pipeline we introduced binaries
which had been linked to muslc by mistake. With this change applied all binaries
will be properly static linked again.

https://github.com/promhippie/prometheus-hetzner-sd/issues/11
