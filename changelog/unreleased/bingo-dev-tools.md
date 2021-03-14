Change: Use bingo for development tooling

We switched to use [bingo](github.com/bwplotka/bingo) for fetching development
and build tools based on fixed defined versions to reduce the dependencies
listed within the regular go.mod file within this project.

https://github.com/promhippie/prometheus-hetzner-sd/issues/14
