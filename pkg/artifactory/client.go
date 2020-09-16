package artifactory

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Client represents an active artifactory object
type Client struct {
	url           string
	restCli       *resty.Client
	systemVersion *SystemVersion
}

// Options holds artifactory client object options
type Options struct {
	URL     string
	APIKey  string
	Timeout time.Duration
}

// New initializes a new artifactory client
func New(opts Options) (*Client, error) {
	var err error

	cli := &Client{
		url: opts.URL,
		restCli: resty.New().
			SetTimeout(opts.Timeout).
			SetHeader("X-JFrog-Art-Api", opts.APIKey),
	}

	cli.systemVersion, err = cli.SystemVersion()
	if err != nil {
		return nil, errors.Wrap(err, "Cannot retrieve Artifactory system version")
	}

	log.Info().Msgf("Artifactory %s (rev. %s) found", cli.systemVersion.Version, cli.systemVersion.Revision)
	return cli, nil
}
