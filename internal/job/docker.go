package job

import (
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/crazy-max/artifactory-cleanup/pkg/artifactory"
	"github.com/docker/go-units"
)

func (j *Job) cleanupDockerRepo(config artifactory.RepoConfiguration) *CleanupResult {
	sublog := j.Log.With().Bool("dry_run", j.DryRun).Str("registry", config.Key).Logger()

	stats := CleanupResult{
		SizeGain:     0,
		ItemsRemoved: 0,
	}

	// Get docker repositories from registry
	repos, err := j.Atf.DockerRepositories(config.Key)
	if err != nil {
		sublog.Error().Err(err).Msg("Cannot retrieve Docker repositories")
	} else if len(repos.Repositories) <= 0 {
		sublog.Warn().Msgf("No Docker repository found on %s", config.Key)
		return nil
	}

	sublog.Info().Msgf("%d repositories found", len(repos.Repositories))

	// Iterate over repositories
	for _, dockerRepo := range repos.Repositories {
		dockerRepoLog := sublog.With().Str("docker_repo", dockerRepo).Logger()

		// Grab tags
		tags, err := j.Atf.DockerTags(config.Key, dockerRepo)
		if err != nil {
			dockerRepoLog.Error().Err(err).Msg("Cannot retrieve Docker repositories")
		} else if len(tags.Tags) <= 0 {
			dockerRepoLog.Warn().Msgf("No Docker tags found")
			continue
		}
		dockerRepoLog.Info().Msgf("%d tags found", len(tags.Tags))

		repoStats := CleanupResult{
			SizeGain:     0,
			ItemsRemoved: 0,
		}

		// Iterate over tags
		for index, tag := range tags.Tags {
			tagLog := dockerRepoLog.With().Str("tag", tag).Logger()

			// Image size
			tagSize, err := j.Atf.DockerTagSize(config.Key, dockerRepo, tag)
			if err != nil {
				tagLog.Warn().Err(err).Msg("Cannot retrieve image size")
			}
			tagLog = tagLog.With().Str("size", units.HumanSize(float64(tagSize))).Logger()

			// Check exclude
			if j.Policy.Docker.IsExcluded(tag) {
				tagLog.Debug().Msg("This tag is excluded. Skipping...")
				continue
			}

			// Check semver
			if *j.Policy.Docker.KeepSemver {
				version, err := semver.NewVersion(tag)
				if err == nil {
					tagLog.Debug().Interface("version", version).Msg("This tag is a valid semver. Skipping...")
					continue
				}
			}

			lastModified, err := j.Atf.DockerTagLastModified(config.Key, dockerRepo, tag)
			if err != nil {
				tagLog.Error().Err(err).Msg("Cannot retrieve last modified date")
				continue
			}

			lastDownloaded, err := j.Atf.DockerTagLastDownloaded(config.Key, dockerRepo, tag)
			if err != nil {
				tagLog.Error().Err(err).Msg("Cannot retrieve last downloaded date")
				continue
			}

			tagDateLog := tagLog.With().
				Time("last_modified", lastModified).
				Time("last_downloaded", lastDownloaded).
				Logger()

			currentDate := time.Now()
			if *j.Policy.LastModified && currentDate.Sub(lastModified) < *j.Policy.Retention {
				tagDateLog.Debug().Msg("Tag last modified is not old enough to be deleted. Skipping...")
				continue
			}
			if *j.Policy.LastDownloaded && currentDate.Sub(lastDownloaded) < *j.Policy.Retention {
				tagDateLog.Debug().Msg("This tag is not old enough to be deleted. Skipping...")
				continue
			}

			if *j.Policy.RetentionCount > 0 && len(tags.Tags)-index <= *j.Policy.RetentionCount {
				tagDateLog.Debug().Msg("This image is within the specified retention count for this repository. Skipping...")
				continue
			}

			if j.DryRun {
				repoStats.SizeGain += tagSize
				repoStats.ItemsRemoved++
				tagDateLog.Info().Msg("Image will be removed")
				continue
			}

			if err := j.Atf.DockerTagRemove(config.Key, dockerRepo, tag); err != nil {
				tagDateLog.Error().Err(err).Msg("Cannot remove image")
			} else {
				repoStats.SizeGain += tagSize
				repoStats.ItemsRemoved++
				tagDateLog.Info().Msg("Image successfully removed")
			}
		}

		stats.SizeGain += repoStats.SizeGain
		stats.ItemsRemoved += repoStats.ItemsRemoved
	}

	return &stats
}
