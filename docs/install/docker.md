# Installation with Docker

## About

Artifactory Cleanup provides automatically updated Docker :whale: images within several registries:

| Registry                                                                                         | Image                           |
|--------------------------------------------------------------------------------------------------|---------------------------------|
| [Docker Hub](https://hub.docker.com/r/crazymax/artifactory-cleanup/)                             | `crazymax/artifactory-cleanup`                 |
| [GitHub Container Registry](https://github.com/users/crazy-max/packages/container/package/artifactory-cleanup)  | `ghcr.io/crazy-max/artifactory-cleanup`        |

It is possible to always use the latest stable tag or to use another service that handles updating Docker images.

!!! note
    Want to be notified of new releases? Check out :bell: [Diun (Docker Image Update Notifier)](https://github.com/crazy-max/diun) project!

Following platforms for this image are available:

```shell
$ docker run --rm mplatform/mquery crazymax/artifactory-cleanup:latest
Image: crazymax/artifactory-cleanup:latest
 * Manifest List: Yes
 * Supported platforms:
   - linux/amd64
   - linux/arm/v6
   - linux/arm/v7
   - linux/arm64
   - linux/386
   - linux/ppc64le
```

This reference setup guides users through the setup based on `docker-compose`, but the installation of `docker-compose`
is out of scope of this documentation. To install `docker-compose` itself, follow the official
[install instructions](https://docs.docker.com/compose/install/).

## Usage

```yaml
version: "3.5"

services:
  artifactory-cleanup:
    image: crazymax/artifactory-cleanup:latest
    container_name: artifactory-cleanup
    environment:
      - "TZ=Europe/Paris"
      - "LOG_LEVEL=info"
      - "LOG_JSON=false"
      - "ATFCLNP_ARTIFACTORY_URL=https://artifactory.example.com"
      - "ATFCLNP_ARTIFACTORY_APIKEY=01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
      - "ATFCLNP_POLICIES_0_NAME=mypolicy"
      - "ATFCLNP_POLICIES_0_REPOS=rpm-prod-local,rpm-local,generic-local"
      - "ATFCLNP_POLICIES_0_SCHEDULE=*/30 * * * *"
      - "ATFCLNP_POLICIES_0_RETENTION=24h"
      - "ATFCLNP_POLICIES_0_LASTMODIFIED=true"
      - "ATFCLNP_POLICIES_0_LASTDOWNLOADED=true"
    restart: always
```

Edit this example with your preferences and run the following commands to bring up Artifactory Cleanup:

```shell
$ docker-compose up -d
$ docker-compose logs -f
```

Or use the following command:

```shell
$ docker run -d --name artifactory-cleanup \
    -e "TZ=Europe/Paris" \
    -e "LOG_LEVEL=info" \
    -e "LOG_JSON=false" \
    -e "ATFCLNP_ARTIFACTORY_URL=https://artifactory.example.com" \
    -e "ATFCLNP_ARTIFACTORY_APIKEY=01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ" \
    -e "ATFCLNP_POLICIES_0_NAME=mypolicy" \
    -e "ATFCLNP_POLICIES_0_REPOS=rpm-prod-local,rpm-local,generic-local" \
    -e "ATFCLNP_POLICIES_0_SCHEDULE=*/30 * * * *" \
    -e "ATFCLNP_POLICIES_0_RETENTION=24h" \
    -e "ATFCLNP_POLICIES_0_LASTMODIFIED=true" \
    -e "ATFCLNP_POLICIES_0_LASTDOWNLOADED=true" \
    crazymax/artifactory-cleanup:latest
```

To upgrade your installation to the latest release:

```shell
$ docker-compose pull
$ docker-compose up -d
```

If you prefer to rely on the [`configuration file](../config/index.md#configuration-file) instead of
environment variables:

```yaml
# ./artifactory-cleanup.yml
artifactory:
  url: "https://artifactory.example.com"
  apiKey: "01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

policies:
  -
    name: "mypolicy"
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

And your docker composition:

```yaml
version: "3.5"

services:
  artifactory-cleanup:
    image: crazymax/artifactory-cleanup:latest
    container_name: artifactory-cleanup
    volumes:
      - "./artifactory-cleanup.yml:/artifactory-cleanup.yml:ro"
    environment:
      - "TZ=Europe/Paris"
      - "LOG_LEVEL=info"
      - "LOG_JSON=false"
    restart: always
```
