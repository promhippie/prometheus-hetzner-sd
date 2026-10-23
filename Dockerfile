FROM --platform=$BUILDPLATFORM golang:1.27.3-alpine@sha256:b8c9b82affbbb82a020bbcf998dad2ba9516944ac7e883524642a481a9505e9c AS builder

RUN apk add --no-cache -U git curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/prometheus-hetzner-sd
COPY . /go/src/prometheus-hetzner-sd/

RUN --mount=type=cache,target=/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task generate build GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.24@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a

RUN apk add --no-cache ca-certificates mailcap && \
    addgroup -g 1337 prometheus-hetzner-sd && \
    adduser -D -u 1337 -h /var/lib/prometheus-hetzner-sd -G prometheus-hetzner-sd prometheus-hetzner-sd

EXPOSE 9000
VOLUME ["/var/lib/prometheus-hetzner-sd"]
ENTRYPOINT ["/usr/bin/prometheus-hetzner-sd"]
CMD ["server"]
HEALTHCHECK CMD ["/usr/bin/prometheus-hetzner-sd", "health"]

ENV PROMETHEUS_HETZNER_OUTPUT_ENGINE="http"
ENV PROMETHEUS_HETZNER_OUTPUT_FILE="/var/lib/prometheus-hetzner-sd/output.json"

COPY --from=builder /go/src/prometheus-hetzner-sd/bin/prometheus-hetzner-sd /usr/bin/prometheus-hetzner-sd
WORKDIR /var/lib/prometheus-hetzner-sd
USER prometheus-hetzner-sd
