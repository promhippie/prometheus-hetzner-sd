#!/bin/sh
set -e

if [ ! -d /var/lib/prometheus-hetzner-sd ] && [ ! -d /etc/prometheus-hetzner-sd ]; then
    userdel prometheus-hetzner-sd 2>/dev/null || true
    groupdel prometheus-hetzner-sd 2>/dev/null || true
fi
