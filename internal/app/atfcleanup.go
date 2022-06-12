package app

import (
	"github.com/crazy-max/artifactory-cleanup/internal/config"
	"github.com/crazy-max/artifactory-cleanup/internal/job"
	"github.com/crazy-max/artifactory-cleanup/pkg/artifactory"
	"github.com/crazy-max/artifactory-cleanup/pkg/utl"
	"github.com/pkg/errors"
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
func New(cfg *config.Config) (*AtfCleanup, error) {
	apiKey, err := utl.GetSecret(cfg.Artifactory.APIKey, cfg.Artifactory.APIKeyFile)
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
		cron: cron.New(cron.WithParser(cron.NewParser(
			cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
		))),
	}, nil
}

// Start starts Artifactory cleanup
func (ac *AtfCleanup) Start() (err error) {
	for _, policy := range ac.cfg.Policies {
		tjob := &job.Job{
			DryRun: ac.cfg.Cli.DryRun,
			Policy: policy,
			Atf:    ac.atf,
			Cron:   ac.cron,
			Log:    log.With().Str("policy", policy.Name).Logger(),
		}

		if ac.cfg.Cli.DisableSchedule {
			tjob.Run()
			continue
		}

		tjob.ID, err = ac.cron.AddJob(policy.Schedule, tjob)
		if err != nil {
			return errors.Wrap(err, "Cannot create job")
		}

		tjob.Log.Info().Msgf("Cron initialized with schedule %s", policy.Schedule)
	}

	if ac.cfg.Cli.DisableSchedule {
		return
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
