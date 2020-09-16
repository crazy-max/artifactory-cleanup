package app

import (
	"time"

	"github.com/crazy-max/artifactory-cleanup/internal/config"
	"github.com/crazy-max/artifactory-cleanup/internal/job"
	"github.com/crazy-max/artifactory-cleanup/pkg/artifactory"
	"github.com/crazy-max/artifactory-cleanup/pkg/utl"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

// AtfCleanup represents an active Artifactory cleanup object
type AtfCleanup struct {
	cfg  *config.Config
	atf  *artifactory.Client
	cron *cron.Cron
}

// New creates a new Artifactory cleanup instance
func New(cfg *config.Config, location *time.Location) (*AtfCleanup, error) {
	apiKey, err := utl.GetSecret(cfg.Artifactory.ApiKey, cfg.Artifactory.ApiKeyFile)
	if err != nil {
		log.Warn().Err(err).Msg("Cannot retrieve API key secret")
	}

	// Artifactory client
	atf, err := artifactory.New(artifactory.Options{
		URL:     cfg.Artifactory.URL,
		APIKey:  apiKey,
		Timeout: *cfg.Artifactory.Timeout,
	})
	if err != nil {
		return nil, err
	}

	return &AtfCleanup{
		cfg: cfg,
		atf: atf,
		cron: cron.New(
			cron.WithLocation(location),
			cron.WithParser(cron.NewParser(cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor)),
		),
	}, nil
}

// Start starts Artifactory cleanup
func (ac *AtfCleanup) Start() error {
	var err error

	for _, policy := range ac.cfg.Policies {
		tjob := &job.Job{
			DryRun: ac.cfg.Cli.DryRun,
			Policy: policy,
			Atf:    ac.atf,
			Cron:   ac.cron,
			Log:    log.With().Str("policy", policy.Name).Logger(),
		}

		tjob.ID, err = ac.cron.AddJob(policy.Schedule, tjob)
		if err != nil {
			tjob.Log.Error().Err(err).Msg("Cannot create job")
			continue
		}

		tjob.Log.Info().Msgf("Cron initialized with schedule %s", policy.Schedule)
	}

	ac.cron.Start()
	select {}
}

// Close closes Artifactory cleanup
func (ac *AtfCleanup) Close() {
	if ac.cron != nil {
		ac.cron.Stop()
	}
}
