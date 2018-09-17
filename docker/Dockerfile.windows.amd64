# escape=`
FROM microsoft/nanoserver:10.0.14393.2430

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" `
  org.label-schema.name="Prometheus Hetzner SD" `
  org.label-schema.vendor="Thomas Boerger" `
  org.label-schema.schema-version="1.0"

ENTRYPOINT ["c:\\prometheus-hetzner-sd.exe"]
CMD ["server"]

COPY bin/prometheus-hetzner-sd.exe c:\prometheus-hetzner-sd.exe
