# Basic example

In this section we quickly go over a basic way to run Artifactory Cleanup.

## Setup

!!! warning
    Make sure to follow the instructions to [install from binary](../install/binary.md) before.

First create a [`artifactory-cleanup.yml` configuration](../config/index.md) file like this one:

```yaml
# ./artifactory-cleanup.yml
artifactory:
  url: "https://artifactory.example.com"
  apiKey: "01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

policies:
  -
    name: "policy_docker"
    repos:
      - "docker-prod-local"
    schedule: "*/30 * * * *"
    retention: "2160h" # 90d
    lastModified: true
    lastDownloaded: true
    docker:
      keepSemver: true
      exclude:
        - "latest"
  -
    name: "policy_misc"
    repos:
      - "rpm-prod-local"
      - "rpm-local"
      - "generic-local"
    schedule: "*/30 * * * *"
    retention: "24h" # 1d
    lastModified: true
    lastDownloaded: true
    common:
      include:
        - "prod/*"
      exclude:
        - "*2.2.*"
        - "*2.1.0*"
```

That's it. Now you can launch Artifactory Cleanup with the following command:

```shell
$ artifactory-cleanup --config ./artifactory-cleanup.yml --dry-run
```
