package artifactory

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// SystemVersion holds system version response
type SystemVersion struct {
	Version  string   `json:"version"`
	Revision string   `json:"revision"`
	Addons   []string `json:"addons"`
	License  string   `json:"license"`
}

// SystemVersion retrieves information about the current Artifactory
// version, revision, and currently installed Add-ons.
func (c *Client) SystemVersion() (*SystemVersion, error) {
	var systemVersion *SystemVersion

	resp, err := c.restCli.R().
		SetHeader("Content-Type", "application/vnd.org.jfrog.artifactory.system.Version+json").
		SetResult(&systemVersion).
		Get(fmt.Sprintf("%s/api/system/version", c.url))
	log.Trace().Interface("response", systemVersion).Msg("Artifactory system version")

	if err != nil {
		return nil, err
	} else if resp.IsError() {
		return nil, errors.New(resp.Status())
	}

	return systemVersion, nil
}
