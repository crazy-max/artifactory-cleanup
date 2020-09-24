package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/alecthomas/kong"
	"github.com/crazy-max/artifactory-cleanup/internal/app"
	"github.com/crazy-max/artifactory-cleanup/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	ac      *app.AtfCleanup
	cli     config.Cli
	name    = "artifactory-cleanup"
	version = "dev"
	meta    = config.Meta{
		ID:     "artifactory-cleanup",
		Name:   "Artifactory Cleanup",
		Desc:   "Cleanup artifacts on Jfrog Artifactory with advanced settings",
		URL:    "https://github.com/crazy-max/artifactory-cleanup",
		Logo:   "https://raw.githubusercontent.com/crazy-max/artifactory-cleanup/master/.github/artifactory-cleanup.png",
		Author: "CrazyMax",
	}
)

func main() {
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())

	meta.Version = version
	meta.UserAgent = fmt.Sprintf("%s/%s go/%s %s", meta.ID, meta.Version, runtime.Version()[2:], strings.Title(runtime.GOOS))

	// Parse command line
	_ = kong.Parse(&cli,
		kong.Name(meta.ID),
		kong.Description(fmt.Sprintf("%s. More info: %s", meta.Desc, meta.URL)),
		kong.UsageOnError(),
		kong.Vars{
			"version": fmt.Sprintf("%s", version),
		},
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))

	// Load timezone location
	location, err := time.LoadLocation(cli.Timezone)
	if err != nil {
		log.Panic().Err(err).Msgf("Cannot load timezone %s", cli.Timezone)
	}

	// Logging
	var logw io.Writer

	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(location)
	}

	if !cli.LogJSON {
		logw = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC1123,
		}).With().Timestamp().Logger()
	} else {
		logw = os.Stdout
	}

	ctx := zerolog.New(logw).With().Timestamp()
	if cli.LogCaller {
		ctx = ctx.Caller()
	}
	log.Logger = ctx.Logger()

	logLevel, err := zerolog.ParseLevel(cli.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msgf("Unknown log level")
	} else {
		zerolog.SetGlobalLevel(logLevel)
	}

	// Init
	log.Info().Str("version", version).Msgf("Starting %s", meta.Name)
	if cli.DryRun {
		log.Warn().Msg("Dry run enabled")
	}

	// Handle os signals
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-channel
		ac.Close()
		log.Warn().Msgf("Caught signal %v", sig)
		os.Exit(0)
	}()

	// Load configuration
	cfg, err := config.Load(cli, meta)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load configuration")
	}
	log.Debug().Msg(cfg.String())

	// Init
	if ac, err = app.New(cfg, location); err != nil {
		log.Fatal().Err(err).Msgf("Cannot initialize %s", meta.Name)
	}

	// Start
	if err = ac.Start(); err != nil {
		log.Fatal().Err(err).Msgf("Cannot start %s", meta.Name)
	}
}
