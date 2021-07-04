# Installation from binary

## Download

Artifactory Cleanup binaries are available on [releases]({{ config.repo_url }}releases/latest) page.

Choose the archive matching the destination platform:

* [`artifactory-cleanup_{{ git.tag | trim('v') }}_darwin_amd64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_darwin_amd64.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_darwin_arm64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_darwin_arm64.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_freebsd_386.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_freebsd_386.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_freebsd_amd64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_freebsd_amd64.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_linux_386.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_linux_386.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_linux_amd64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_linux_amd64.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_linux_arm64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_linux_arm64.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_linux_armv5.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_linux_armv5.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_linux_armv6.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_linux_armv6.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_linux_armv7.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_linux_armv7.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_linux_ppc64le.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_linux_ppc64le.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_linux_riscv64.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_linux_riscv64.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_linux_s390x.tar.gz`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_linux_s390x.tar.gz)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_windows_386.zip`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_windows_386.zip)
* [`artifactory-cleanup_{{ git.tag | trim('v') }}_windows_amd64.zip`]({{ config.repo_url }}/releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_windows_amd64.zip)

And extract Artifactory Cleanup:

```shell
wget -qO- {{ config.repo_url }}releases/download/v{{ git.tag | trim('v') }}/artifactory-cleanup_{{ git.tag | trim('v') }}_linux_amd64.tar.gz | tar -zxvf - artifactory-cleanup
```

After getting the binary, it can be tested with [`./artifactory-cleanup --help`](../usage/cli.md) command and moved
to a permanent location.

## Server configuration

Steps below are the recommended server configuration.

### Prepare environment

Create user to run Artifactory Cleanup (ex. `artifactory-cleanup`)

```shell
groupadd artifactory-cleanup
useradd -s /bin/false -d /bin/null -g artifactory-cleanup artifactory-cleanup
```

### Create required directory structure

```shell
mkdir -p /var/lib/artifactory-cleanup
chown artifactory-cleanup:artifactory-cleanup /var/lib/artifactory-cleanup/
chmod -R 750 /var/lib/artifactory-cleanup/
mkdir /etc/artifactory-cleanup
chown artifactory-cleanup:artifactory-cleanup /etc/artifactory-cleanup
chmod 770 /etc/artifactory-cleanup
```

### Configuration

Create your first [configuration](../config/index.md) file in `/etc/artifactory-cleanup/artifactory-cleanup.yml`
and type:

```shell
chown artifactory-cleanup:artifactory-cleanup /etc/artifactory-cleanup/artifactory-cleanup.yml
chmod 644 /etc/artifactory-cleanup/artifactory-cleanup.yml
```

### Copy binary to global location

```shell
cp artifactory-cleanup /usr/local/bin/artifactory-cleanup
```

## Running Artifactory Cleanup

After the above steps, two options to run Artifactory Cleanup:

### 1. Creating a service file (recommended)

See how to create [Linux service](linux-service.md) to start Artifactory Cleanup automatically.

### 2. Running from terminal

```shell
/usr/local/bin/artifactory-cleanup \
  --config /etc/artifactory-cleanup/artifactory-cleanup.yml \
  --dry-run
```

## Updating to a new version

You can update to a new version of Artifactory Cleanup by stopping it, replacing the binary
at `/usr/local/bin/artifactory-cleanup` and restarting the instance.

If you have carried out the installation steps as described above, the binary should have the generic
name `artifactory-cleanup`. Do not change this, i.e. to include the version number.
