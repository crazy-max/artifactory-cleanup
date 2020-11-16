# Command Line

## Usage

```shell
$ artifactory-cleanup [options]
```

## Options

```
$ artifactory-cleanup --help
Usage: artifactory-cleanup

Cleanup artifacts on Jfrog Artifactory with advanced settings. More info:
https://github.com/crazy-max/artifactory-cleanup

Flags:
  -h, --help                Show context-sensitive help.
      --version
      --config=STRING       Configuration file ($CONFIG).
      --log-level="info"    Set log level ($LOG_LEVEL).
      --log-json            Enable JSON logging output ($LOG_JSON).
      --log-caller          Add file:line of the caller to log output
                            ($LOG_CALLER).
      --dry-run             If enabled files will not be removed ($DRY_RUN).
      --disable-schedule    Disable scheduling and execute policies right away
                            ($DISABLE_SCHEDULE).
```

## Environment variables

Following environment variables can be used in place:

| Name               | Default       | Description   |
|--------------------|---------------|---------------|
| `CONFIG`           |               | Configuration file |
| `LOG_LEVEL`        | `info`        | Log level output |
| `LOG_JSON`         | `false`       | Enable JSON logging output |
| `LOG_CALLER`       | `false`       | Enable to add `file:line` of the caller |
| `DRY_RUN`          | `false`       | If enabled files will not be removed |
| `DISABLE_SCHEDULE` | `false`       | Disable scheduling and execute policies right away |
