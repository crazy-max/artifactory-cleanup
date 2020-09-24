package config

import (
	"github.com/alecthomas/kong"
)

// Cli holds command line args, flags and cmds
type Cli struct {
	Version         kong.VersionFlag
	Cfgfile         string `kong:"name='config',env='CONFIG',help='Configuration file.'"`
	Timezone        string `kong:"name='timezone',env='TZ',default='UTC',help='Timezone.'"`
	LogLevel        string `kong:"name='log-level',env='LOG_LEVEL',default='info',help='Set log level.'"`
	LogJSON         bool   `kong:"name='log-json',env='LOG_JSON',default='false',help='Enable JSON logging output.'"`
	LogCaller       bool   `kong:"name='log-caller',env='LOG_CALLER',default='false',help='Add file:line of the caller to log output.'"`
	DryRun          bool   `kong:"name='dry-run',env='DRY_RUN',default='false',help='If enabled files will not be removed.'"`
	DisableSchedule bool   `kong:"name='disable-schedule',env='DISABLE_SCHEDULE',default='false',help='Disable scheduling and execute policies right away.'"`
}
