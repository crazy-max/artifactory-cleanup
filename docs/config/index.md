# Configuration

## Overview

There are two different ways to define configuration in Artifactory Cleanup:

* In a [configuration file](#configuration-file)
* As [environment variables](#environment-variables)

These ways are evaluated in the order listed above.

If no value was provided for a given option, a default value applies. Moreover, if an option has sub-options, and
any of these sub-options is not specified, a default value will apply as well.

## Configuration file

At startup, Artifactory Cleanup searches for a file named `artifactory-cleanup.yml` (or `artifactory-cleanup.yaml`) in:

* `/etc/artifactory-cleanup/`
* `$XDG_CONFIG_HOME/`
* `$HOME/.config/`
* `.` _(the working directory)_

You can override this using the [`--config` flag or `CONFIG` env var](../usage/cli.md).

??? example "artifactory-cleanup.yml"
    ```yaml
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

## Environment variables

All configuration from file can be transposed into environment variables. As an example, the following configuration:

??? example "artifactory-cleanup.yml"
    ```yaml
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

Can be transposed to:

??? example "environment variables"
    ```
    ATFCLNP_ARTIFACTORY_URL=https://artifactory.example.com
    ATFCLNP_ARTIFACTORY_APIKEY=01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ
    
    ATFCLNP_POLICIES_0_NAME=mypolicy
    ATFCLNP_POLICIES_0_REPOS=rpm-prod-local,rpm-local,generic-local
    ATFCLNP_POLICIES_0_SCHEDULE=*/30 * * * *
    ATFCLNP_POLICIES_0_RETENTION=24h
    ATFCLNP_POLICIES_0_LASTMODIFIED=true
    ATFCLNP_POLICIES_0_LASTDOWNLOADED=true
    ```

## Reference

* [artifactory](artifactory.md)
* [policies](policies.md)
