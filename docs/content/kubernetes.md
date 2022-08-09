---
title: "Kubernetes"
date: 2022-07-22T00:00:00+00:00
anchor: "kubernetes"
weight: 20
---

## Kubernetes

Currently we are covering the most famous installation methods on Kubernetes,
you can choose between [Kustomize][kustomize] and [Helm][helm].

### Kustomize

We won't cover the installation of [Kustomize][kustomize] within this guide, to
get it installed and working please read the upstream documentation. After the
installation of [Kustomize][kustomize] you just need to prepare a
`kustomization.yml` wherever you like similar to this:

{{< highlight yaml >}}
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: prometheus-hetzner-sd

resources:
  - github.com/promhippie/prometheus-hetzner-sd//deploy/kubernetes?ref=master

configMapGenerator:
  - name: prometheus-hetzner-sd
    behavior: merge
    literals: []

secretGenerator:
  - name: prometheus-hetzner-sd
    behavior: merge
    literals: []
{{< / highlight >}}

After that you can simply execute `kustomize build | kubectl apply -f -` to get
the manifest applied. Generally it's best to use fixed versions of the container
images, this can be done quite easy, you just need to append this block to your
`kustomization.yml` to use this specific version:

{{< highlight yaml >}}
images:
  - name: quay.io/promhippie/prometheus-hetzner-sd
    newTag: 1.1.0
{{< / highlight >}}

After applying this manifest the exporter should be directly visible within your
Prometheus instance if you are using the Prometheus Operator as these manifests
are providing a ServiceMonitor.

To consume the service discovery within Prometheus you got to configre matching
scrape targets using the HTTP engine, just add a block similar to this one to
your Prometheus configuration:

{{< highlight yaml >}}
scrape_configs:
- job_name: node
  http_sd_configs:
  - url: http://hetzner-sd.prometheus-hetzner-sd.svc.cluster.local:9000/sd
  relabel_configs:
  - source_labels: [__meta_hetzner_public_ipv4]
    replacement: "${1}:9100"
    target_label: __address__
  - source_labels: [__meta_hetzner_dc]
    target_label: datacenter
  - source_labels: [__meta_hetzner_name]
    target_label: instance
{{< / highlight >}}

### Helm

We won't cover the installation of [Helm][helm] within this guide, to get it
installed and working please read the upstream documentation. After the
installation of [Helm][helm] you just need to execute the following commands:

{{< highlight console >}}
helm repo add promhippie https://promhippie.github.io/charts
helm show values promhippie/prometheus-hetzner-sd
helm install prometheus-hetzner-sd promhippie/prometheus-hetzner-sd
{{< / highlight >}}

You can also watch that available values and generally the details of the chart
provided by us within our [chart][chart] repository.

After applying this manifest the exporter should be directly visible within your
Prometheus instance depending on your installation if you enabled the
annotations or the service monitor.

To consume the service discovery within Prometheus you got to configre matching
scrape targets using the HTTP engine, just add a block similar to this one to
your Prometheus configuration:

{{< highlight yaml >}}
scrape_configs:
- job_name: node
  http_sd_configs:
  - url: http://hetzner-sd.prometheus-hetzner-sd.svc.cluster.local:9000/sd
  relabel_configs:
  - source_labels: [__meta_hetzner_public_ipv4]
    replacement: "${1}:9100"
    target_label: __address__
  - source_labels: [__meta_hetzner_dc]
    target_label: datacenter
  - source_labels: [__meta_hetzner_name]
    target_label: instance
{{< / highlight >}}

[kustomize]: https://github.com/kubernetes-sigs/kustomize
[helm]: https://helm.sh
[chart]: https://github.com/promhippie/charts/tree/master/charts/prometheus-hetzner-sd
