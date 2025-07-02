#!/bin/sh
set -e

chown -R prometheus-hetzner-sd:prometheus-hetzner-sd /etc/prometheus-hetzner-sd
chown -R prometheus-hetzner-sd:prometheus-hetzner-sd /var/lib/prometheus-hetzner-sd
chmod 750 /var/lib/prometheus-hetzner-sd

if [ -d /run/systemd/system ]; then
    systemctl daemon-reload

    if systemctl is-enabled --quiet prometheus-hetzner-sd.service; then
        systemctl restart prometheus-hetzner-sd.service
    fi
fi
