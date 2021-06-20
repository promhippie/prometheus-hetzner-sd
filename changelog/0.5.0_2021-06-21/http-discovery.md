Change: Integrate new HTTP service discovery handler

We integrated the new HTTP service discovery which have been introduced by
Prometheus starting with version 2.28. With this new service discovery you can
deploy this service whereever you want and you are not tied to the Prometheus
filesystem anymore.

https://github.com/promhippie/prometheus-hetzner-sd/issues/35
