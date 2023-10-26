Change: Read secrets form files

We have added proper support to load secrets like passwords from files or from
base64-encoded strings. Just provide the flags or environment variables with a
DSN formatted string like `file://path/to/file` or `base64://Zm9vYmFy`.

https://github.com/promhippie/prometheus-hetzner-sd/issues/205
