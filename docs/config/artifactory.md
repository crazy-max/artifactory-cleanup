# Artifactory configuration

```yaml
artifactory:
  url: "https://artifactory.example.com"
  apiKey: "01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
```

## `url`

Artifcatory URL.

!!! example "Config file"
    ```yaml
    artifactory:
      url: "https://artifactory.example.com"
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_ARTIFACTORY_URL`

## `apiKey`

Artifactory [API key](https://www.jfrog.com/confluence/display/JFROG/User+Profile).

!!! example "Config file"
    ```yaml
    artifactory:
      apiKey: "01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_ARTIFACTORY_APIKEY`

## `apiKeyFile`

Use content of secret file as Artifactory API key if `apiKey` not defined.

!!! example "Config file"
    ```yaml
    credentials:
      apiKeyFile: /run/secrets/apikey
    ```

!!! abstract "Environment variables"
    * `ATFCLNP_ARTIFACTORY_APIKEYFILE`
