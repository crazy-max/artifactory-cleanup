package artifactory

import (
	"fmt"
	"path"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// DockerRepositories holds docker repositories response
type DockerRepositories struct {
	Repositories []string `json:"repositories"`
}

// DockerTags holds docker tags response
type DockerTags struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// DockerRepositories returns a list of docker repositories
func (c *Client) DockerRepositories(registry string) (dr DockerRepositories, err error) {
	resp, err := c.restCli.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&dr).
		Get(fmt.Sprintf("%s/api/docker/%s/v2/_catalog", c.url, registry))

	log.Trace().Interface("repositories", dr).Msg("Artifactory docker repositories response")

	if err != nil {
		return dr, err
	} else if resp.IsError() {
		return dr, errors.New(resp.Status())
	}

	return dr, err
}

// DockerTags returns tags of a docker repository
func (c *Client) DockerTags(registry string, repo string) (dt DockerTags, err error) {
	resp, err := c.restCli.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&dt).
		Get(fmt.Sprintf("%s/api/docker/%s/v2/%s/tags/list", c.url, registry, repo))

	log.Trace().Interface("tags", dt).Msg("Artifactory docker tags response")

	if err != nil {
		return dt, err
	} else if resp.IsError() {
		return dt, errors.New(resp.Status())
	}

	return dt, err
}

// DockerTagSize returns the size of a Docker tag
func (c *Client) DockerTagSize(registry string, repo string, tag string) (int64, error) {
	var imageSize = int64(0)

	fs, err := c.FileList(registry, fmt.Sprintf("%s/%s", repo, tag))
	if err != nil {
		return imageSize, err
	}

	// Iterate over files
	for _, file := range fs.Files {
		imageSize += file.Size
	}

	return imageSize, nil
}

// DockerTagLastModified returns the last modified time of a Docker tag manifest
func (c *Client) DockerTagLastModified(registry string, repo string, tag string) (time.Time, error) {
	var lm LastModified

	resp, err := c.restCli.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&lm).
		Get(fmt.Sprintf("%s/api/storage/%s/%s/%s/manifest.json?lastModified", c.url, registry, repo, tag))

	log.Trace().Interface("last_modified", lm).Msg("Artifactory docker tag last modified response")

	if err != nil {
		return time.Time{}, err
	} else if resp.IsError() {
		return time.Time{}, errors.New(resp.Status())
	}

	return time.Parse("2006-01-02T15:04:05.000-0700", lm.LastModified)
}

// DockerTagLastDownloaded returns the last downloaded time of a Docker tag manifest
func (c *Client) DockerTagLastDownloaded(registry string, repo string, tag string) (time.Time, error) {
	return c.LastDownloaded(registry, path.Join(repo, tag))
}

// DockerTagRemove removes a Docker tag on Artifactory
func (c *Client) DockerTagRemove(registry string, repo string, tag string) error {
	return c.DeleteItem(registry, fmt.Sprintf("%s/%s", repo, tag))
}
