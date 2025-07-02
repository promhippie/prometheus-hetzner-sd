#!/bin/sh
set -e

if ! getent group prometheus-hetzner-sd >/dev/null 2>&1; then
    groupadd --system prometheus-hetzner-sd
fi

if ! getent passwd prometheus-hetzner-sd >/dev/null 2>&1; then
    useradd --system --create-home --home-dir /var/lib/prometheus-hetzner-sd --shell /bin/bash -g prometheus-hetzner-sd prometheus-hetzner-sd
fi
