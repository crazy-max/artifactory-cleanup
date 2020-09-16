package artifactory

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// RepoConfiguration holds repository configuration response
type RepoConfiguration struct {
	Key         string    `json:"key"`
	Rclass      RepoClass `json:"rclass"`
	PackageType string    `json:"packageType"`
	Description string    `json:"description"`
	Notes       string    `json:"notes"`
}

// RepoClass constants
const (
	RepoClassLocal   = RepoClass("local")
	RepoClassRemote  = RepoClass("remote")
	RepoClassVirtual = RepoClass("virtual")
)

// RepoClass holds repo classification type
type RepoClass string

// RepoConfiguration retrieves the current configuration of a repository
func (c *Client) RepoConfiguration(repo string) (RepoConfiguration, error) {
	var rc RepoConfiguration

	resp, err := c.restCli.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&rc).
		Get(fmt.Sprintf("%s/api/repositories/%s", c.url, repo))

	log.Trace().Interface("repo_configuration", rc).Msg("Artifactory repo configuration response")

	if err != nil {
		return rc, err
	} else if resp.IsError() {
		return rc, errors.New(resp.Status())
	}

	return rc, err
}
