package job

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/crazy-max/artifactory-cleanup/pkg/artifactory"
	"github.com/docker/go-units"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
)

func (j *Job) cleanupCommonRepo(config artifactory.RepoConfiguration) *CleanupResult {
	var files []artifactory.AQLResult
	sublog := j.Log.With().Bool("dry_run", j.DryRun).Str("repo", config.Key).Logger()

	stats := CleanupResult{
		SizeGain:     0,
		ItemsRemoved: 0,
	}

	for _, include := range j.Policy.Common.Include {
		aqlQuery, err := utils.CreateAqlBodyForSpecWithPattern(&utils.CommonParams{
			Pattern:     fmt.Sprintf("%s/%s", config.Key, strings.TrimLeft(include, "/")),
			Exclusions:  j.Policy.Common.Exclude,
			Offset:      0,
			Limit:       0,
			Recursive:   true,
			IncludeDirs: false,
			Regexp:      false,
		})
		if err != nil {
			sublog.Error().Err(err).Msgf("Cannot create AQL query for include pattern: %s", include)
			continue
		}

		res, err := j.Atf.SearchAQL(aqlQuery)
		if err != nil {
			sublog.Error().Err(err).Msg("Cannot find files")
			continue
		}

		files = uniqueAQLResult(append(files, res.Results...))
	}

	sublog.Info().Msgf("%d files found", len(files))

	// Iterate over files
	for index, file := range files {
		filelog := sublog.With().
			Str("file", path.Join(file.Path, file.Name)).
			Str("size", units.HumanSize(float64(file.Size))).
			Time("last_modified", file.Modified).Logger()

		lastDownloaded, err := j.Atf.LastDownloaded(file.Repo, path.Join(file.Path, file.Name))
		if err != nil {
			filelog.Error().Err(err).Msg("Cannot retrieve last downloaded date")
			continue
		}

		filelog = filelog.With().
			Time("last_downloaded", lastDownloaded).
			Logger()

		currentDate := time.Now()
		if *j.Policy.LastModified && currentDate.Sub(file.Modified) < *j.Policy.Retention {
			filelog.Debug().Msg("File last modified is not old enough to be deleted. Skipping...")
			continue
		}
		if *j.Policy.LastDownloaded && currentDate.Sub(lastDownloaded) < *j.Policy.Retention {
			filelog.Debug().Msg("This file is not old enough to be deleted. Skipping...")
			continue
		}

		if *j.Policy.RetentionCount > 0 && len(files)-index <= *j.Policy.RetentionCount {
			filelog.Debug().Msg("This image is within the specified retention count for this repository. Skipping...")
			continue
		}

		if j.DryRun {
			stats.SizeGain += file.Size
			stats.ItemsRemoved++
			filelog.Info().Msg("File will be removed")
			continue
		}

		if err := j.Atf.DeleteItem(config.Key, path.Join(file.Path, file.Name)); err != nil {
			filelog.Error().Err(err).Msg("Cannot remove file")
		} else {
			stats.SizeGain += file.Size
			stats.ItemsRemoved++
			filelog.Info().Msg("File successfully removed")
		}
	}

	return &stats
}

func uniqueAQLResult(aqlResult []artifactory.AQLResult) []artifactory.AQLResult {
	var unique []artifactory.AQLResult
aqlResultLoop:
	for _, v := range aqlResult {
		for i, u := range unique {
			if v.Repo == u.Repo && v.Path == u.Path && v.Name == u.Name {
				unique[i] = v
				continue aqlResultLoop
			}
		}
		unique = append(unique, v)
	}
	return unique
}
