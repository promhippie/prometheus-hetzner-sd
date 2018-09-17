FROM webhippie/alpine:latest AS build
RUN apk add --no-cache ca-certificates mailcap

FROM scratch

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Prometheus Hetzner SD" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

ENTRYPOINT ["/usr/bin/prometheus-hetzner-sd"]
CMD ["server"]

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/mime.types /etc/

COPY dist/binaries/prometheus-hetzner-sd-*-linux-arm64 /usr/bin/prometheus-hetzner-sd
