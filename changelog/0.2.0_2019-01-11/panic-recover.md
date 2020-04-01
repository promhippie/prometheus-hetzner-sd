Change: Panic recover within handlers

To make sure panics are properly handled we added a middleware to recover
properly from panics.

https://github.com/promhippie/prometheus-hetzner-sd/pull/3
