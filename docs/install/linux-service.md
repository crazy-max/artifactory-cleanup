# Run as service on Debian based distro

## Using systemd

!!! warning
    Make sure to follow the instructions to [install from binary](binary.md) before.

To create a new service, paste this content in `/etc/systemd/system/artifactory-cleanup.service`:

```
[Unit]
Description=artifactory-cleanup
Documentation={{ config.site_url }}
After=syslog.target
After=network.target

[Service]
RestartSec=2s
Type=simple
User=artifactory-cleanup
Group=artifactory-cleanup
ExecStart=/usr/local/bin/artifactory-cleanup --config /etc/artifactory-cleanup/artifactory-cleanup.yml
Restart=always
#Environment=TZ=Europe/Paris

[Install]
WantedBy=multi-user.target
```

Change the user, group, and other required startup values following your needs.

Enable and start artifactory-cleanup at boot:

```shell
sudo systemctl enable artifactory-cleanup
sudo systemctl start artifactory-cleanup
```

To view logs:

```shell
journalctl -fu artifactory-cleanup.service
```
