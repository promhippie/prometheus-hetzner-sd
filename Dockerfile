FROM --platform=$BUILDPLATFORM golang:1.26.4-alpine@sha256:f1ddd9fe14fffc091dd98cb4bfa999f32c5fc77d2f2305ea9f0e2595c5437c14 AS builder

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

FROM alpine:3.24@sha256:f5064d3e5f88c467c714509f491853ab2d951932c5cad699c0cb969dcec6f3b4

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
