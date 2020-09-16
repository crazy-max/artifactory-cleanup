package artifactory

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// AQLSearchResults holds AQL search results response
type AQLSearchResults struct {
	Results []AQLResult `json:"results,omitempty"`
}

// AQLResult holds AQL result response
type AQLResult struct {
	Repo       string    `json:"repo,omitempty"`
	Path       string    `json:"path,omitempty"`
	Name       string    `json:"name,omitempty"`
	Type       string    `json:"type,omitempty"`
	Size       int64     `json:"size,omitempty"`
	Created    time.Time `json:"created,omitempty"`
	CreatedBy  string    `json:"created_by,omitempty"`
	Modified   time.Time `json:"modified,omitempty"`
	ModifiedBy string    `json:"modified_by,omitempty"`
	Updated    time.Time `json:"updated,omitempty"`
}

// SearchAQL search files using AQL language
func (c *Client) SearchAQL(query string) (AQLSearchResults, error) {
	var asr AQLSearchResults

	resp, err := c.restCli.R().
		SetResult(&asr).
		SetBody([]byte(fmt.Sprintf(`items.find(%s)`, query))).
		Post(fmt.Sprintf("%s/api/search/aql", c.url))

	log.Trace().Interface("search_aql", asr).Str("query", query).Msg("Artifactory search AQL response")

	if err != nil {
		return asr, err
	} else if resp.IsError() {
		return asr, errors.New(resp.Status())
	}

	return asr, err
}
