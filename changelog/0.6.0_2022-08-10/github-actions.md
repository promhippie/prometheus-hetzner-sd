Enhancement: Use GitHub Actions onstead of Drone CI

We have replaced the previous Drone CI setup by more simple GitHub Actions since
are anyway using GitHub for the code hosting and issue tracking. As part of that
we are now also publishing the docker images to Quay.

https://github.com/promhippie/prometheus-hetzner-sd/pull/121
