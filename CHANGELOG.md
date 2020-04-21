# Changelog for unreleased

The following sections list the changes for unreleased.

## Summary

 * Chg #10: Code and project restructuring

## Details

 * Change #10: Code and project restructuring

   To get the project and code structure into a new shape and to get it cleaned up we switched to Go
   modules and restructured the project source in general. The functionality stays the same as
   before.

   https://github.com/promhippie/prometheus-hetzner-sd/pull/10


# Changelog for 0.3.0

The following sections list the changes for 0.3.0.

## Summary

 * Chg #8: Define healthcheck command
 * Chg #5: Support for multiple accounts
 * Chg #4: Switch to cloud.drone.io

## Details

 * Change #8: Define healthcheck command

   To check the health status of the service discovery especially within Docker we added a simple
   subcommand which checks the healthz endpoint to show if the service is up and running.

   https://github.com/promhippie/prometheus-hetzner-sd/pull/8

 * Change #5: Support for multiple accounts

   Make the deployments of this service discovery easier, previously we had to launch one
   instance for every credentials we wanted to gather, with this change we are able to define
   multiple credentials for a single instance of the service discovery.

   https://github.com/promhippie/prometheus-hetzner-sd/pull/5

 * Change #4: Switch to cloud.drone.io

   We don't wanted to maintain our own Drone infrastructure anymore, since there is
   cloud.drone.io available for free we switched the pipelines over to it.

   https://github.com/promhippie/prometheus-hetzner-sd/pull/4


# Changelog for 0.2.0

The following sections list the changes for 0.2.0.

## Summary

 * Chg #1: Add basic documentation
 * Chg #2: Pin xgo to golang 1.10 to avoid issues
 * Chg #3: Update dependencies
 * Chg #3: Lowercase datacenter label
 * Chg #3: Timeout for metrics handler
 * Chg #3: Panic recover within handlers

## Details

 * Change #1: Add basic documentation

   Add some basic documentation page which also includes build and installation instructions to
   make clear how this project can be installed and used.

   https://github.com/promhippie/prometheus-hetzner-sd/pull/1

 * Change #2: Pin xgo to golang 1.10 to avoid issues

   There had been issues while using the latest xgo version, let's pin this tag to 1.10 to ensure the
   binaries are properly build.

   https://github.com/promhippie/prometheus-hetzner-sd/pull/2

 * Change #3: Update dependencies

   Just make sure to update all the build dependencies to work with the latest versions available.

   https://github.com/promhippie/prometheus-hetzner-sd/pull/3

 * Change #3: Lowercase datacenter label

   To get the datacenter name labels normalized we are simply lowercasing the value from now on.

   https://github.com/promhippie/prometheus-hetzner-sd/pull/3

 * Change #3: Timeout for metrics handler

   We added an additional middleware to properly timeout requests to the metrics endpoint for
   long running request.

   https://github.com/promhippie/prometheus-hetzner-sd/pull/3

 * Change #3: Panic recover within handlers

   To make sure panics are properly handled we added a middleware to recover properly from panics.

   https://github.com/promhippie/prometheus-hetzner-sd/pull/3


# Changelog for 0.1.0

The following sections list the changes for 0.1.0.

## Summary

 * Chg #9: Initial release of basic version

## Details

 * Change #9: Initial release of basic version

   Just prepared an initial basic version which could be released to the public.

   https://github.com/promhippie/prometheus-hetzner-sd/issues/9


