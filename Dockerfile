FROM --platform=$BUILDPLATFORM golang:1.26.3-alpine@sha256:91eda9776261207ea25fd06b5b7fed8d397dd2c0a283e77f2ab6e91bfa71079d AS builder

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

FROM alpine:3.23@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11

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
