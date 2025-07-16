#!/bin/sh
set -e

systemctl stop prometheus-hetzner-sd.service || true
systemctl disable prometheus-hetzner-sd.service || true
