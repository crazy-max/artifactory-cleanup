package job

import (
	"sync/atomic"
	"time"

	"github.com/crazy-max/artifactory-cleanup/internal/config"
	"github.com/crazy-max/artifactory-cleanup/pkg/artifactory"
	"github.com/docker/go-units"
	"github.com/hako/durafmt"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
)

type CleanupResult struct {
	SizeGain     int64
	ItemsRemoved int64
}

// Job holds job object
type Job struct {
	ID     cron.EntryID
	DryRun bool
	Policy config.Policy
	Atf    *artifactory.Client
	Cron   *cron.Cron
	Log    zerolog.Logger

	locker uint32
}

// Run runs the job
func (j *Job) Run() {
	if !atomic.CompareAndSwapUint32(&j.locker, 0, 1) {
		j.Log.Warn().Msg("Already running")
		return
	}
	defer atomic.StoreUint32(&j.locker, 0)

	if j.ID > 0 {
		defer j.Log.Info().Msgf("Next run in %s (%s)",
			durafmt.Parse(time.Until(j.Cron.Entry(j.ID).Next)).LimitFirstN(2).String(),
			j.Cron.Entry(j.ID).Next)
	}

	j.Log.Info().Msg("Job triggered")
	for _, repo := range j.Policy.Repos {
		repoCfg, err := j.Atf.RepoConfiguration(repo)
		if err != nil {
			j.Log.Error().Err(err).Msgf("Cannot retrieve repo configuration for %s", repo)
			continue
		}

		if repoCfg.Rclass != artifactory.RepoClassLocal {
			j.Log.Error().Err(err).Msgf("%s is not a local repository", repo)
			continue
		}

		var res *CleanupResult
		switch repoCfg.PackageType {
		case "docker":
			res = j.cleanupDockerRepo(repoCfg)
		default:
			res = j.cleanupCommonRepo(repoCfg)
		}

		if res != nil {
			j.Log.Info().
				Str("repo", repo).
				Str("size_gain", units.HumanSize(float64(res.SizeGain))).
				Int64("items_removed", res.ItemsRemoved).
				Msg("Cleanup done!")
		}
	}
}
