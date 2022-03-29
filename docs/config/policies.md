# Policies configuration

Slice of policies to use during cleanup job.

```yaml
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
      repoRetentionCount: 3
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

### `name`

Name of the policy.

!!! example "Config file"
    ```yaml
    policies:
      - name: "mypolicy"
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_POLICIES_<KEY>_NAME`

### `repos`

A list of repositories to clean.

!!! example "Config file"
    ```yaml
    policies:
      - name: "mypolicy"
        repos:
          - "rpm-prod-local"
          - "rpm-local"
          - "generic-local"
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_POLICIES_<KEY>_REPOS`

### `schedule`

[CRON expression](https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format) to schedule this policy.

!!! example "Config file"
    ```yaml
    policies:
      - name: "mypolicy"
        schedule: "*/30 * * * *"
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_POLICIES_<KEY>_SCHEDULE`

### `retention`

Interval duration to look back before deleting an artifact.

!!! example "Config file"
    ```yaml
    policies:
      - name: "mypolicy"
        retention: "24h"
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_POLICIES_<KEY>_RETENTION`

### `lastModified`

Use last modified time of an artifact as `retention` duration. (default `true`)

!!! example "Config file"
    ```yaml
    policies:
      - name: "mypolicy"
        lastModified: true
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_POLICIES_<KEY>_LASTMODIFIED`

### `lastDownloaded`

Use last downloaded time of an artifact as `retention` duration. (default `true`)

!!! example "Config file"
    ```yaml
    policies:
      - name: "mypolicy"
        lastDownloaded: true
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_POLICIES_<KEY>_LASTDOWNLOADED`

### `common`

The `common` type will apply to repositories via an [AQL search query](https://www.jfrog.com/confluence/display/JFROG/Artifactory+Query+Language)
with inclusion and exclusion fields. However, it will only concern repository types not currently managed by
Artifactory Cleanup such as `generic`, `maven`, `npm`, `rpm` and so on.

!!! example "Config file"
    ```yaml
    policies:
      -
        name: "mypolicy"
        common:
          include:
            - "prod/*"
          exclude:
            - "*2.2.*"
            - "*2.1.0*"
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_POLICIES_<KEY>_COMMON_INCLUDE`
    * `ATFCLNP_POLICIES_<KEY>_COMMON_EXCLUDE`

| Name               | Default       | Description   |
|--------------------|---------------|---------------|
| `include`          |               | List of items matching an expression to include |
| `exclude`          |               | List of items matching an expression to exclude |

### `docker`

The `docker` type will be only applied to `docker` package type repositories on Artifactory. It will work differently
from the `common` type because of the particular structure of a registry. Indeed only the Docker tags will be analyzed
during the cleanup process.

!!! example "Config file"
    ```yaml
    policies:
      -
        name: "mypolicy"
        docker:
          keepSemver: true
          repoRetentionCount: 3
          exclude:
            - "latest"
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_POLICIES_<KEY>_DOCKER_KEEPSEMVER`
    * `ATFCLNP_POLICIES_<KEY>_DOCKER_EXCLUDE`
    * `ATFCLNP_POLICIES_<KEY>_DOCKER_REPORETENTIONCOUNT`

| Name                 | Default       | Description   |
|----------------------|---------------|---------------|
| `keepSemver`         |               | Do not remove tags matching a semver compliant pattern |
| `repoRetentionCount` |               | The number of images that should be kept in each docker repository regardless of `lastModified` and `lastDownloaded` |
| `exclude`            |               | List of tags matching an expression to exclude |
