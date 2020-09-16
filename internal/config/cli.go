package config

import (
	"github.com/alecthomas/kong"
)

// Cli holds command line args, flags and cmds
type Cli struct {
	Version    kong.VersionFlag
	Cfgfile    string `kong:"name='config',env='CONFIG',help='Configuration file.'"`
	Timezone   string `kong:"name='timezone',default='UTC',help='Timezone.'"`
	LogLevel   string `kong:"name='log-level',default='info',help='Set log level.'"`
	LogJSON    bool   `kong:"name='log-json',default='false',help='Enable JSON logging output.'"`
	LogNoColor bool   `kong:"name='log-no-color',default='false',help='Disable coloring output.'"`
	DryRun     bool   `kong:"name='dry-run',help='If enabled, images will not be removed.'"`
}
