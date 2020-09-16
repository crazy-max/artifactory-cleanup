package config

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/crazy-max/artifactory-cleanup/pkg/utl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadFile(t *testing.T) {
	cases := []struct {
		name     string
		cli      Cli
		wantData *Config
		wantErr  bool
	}{
		{
			name:    "Failed on non-existing file",
			wantErr: true,
		},
		{
			name: "Fail on wrong file format",
			cli: Cli{
				Cfgfile: "./fixtures/config.invalid.yml",
			},
			wantErr: true,
		},
		{
			name: "Success",
			cli: Cli{
				Cfgfile: "./fixtures/config.test.yml",
			},
			wantData: &Config{
				Cli: Cli{
					Cfgfile: "./fixtures/config.test.yml",
				},
				Artifactory: &Artifactory{
					URL:     "https://artifactory.example.com",
					ApiKey:  "01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ",
					Timeout: utl.NewDuration(10 * time.Second),
				},
				Policies: []Policy{
					{
						Name: "policy_docker",
						Repos: []string{
							"docker-me-local",
						},
						Schedule:       "*/30 * * * *",
						Retention:      utl.NewDuration(720 * time.Hour),
						LastModified:   utl.NewTrue(),
						LastDownloaded: utl.NewTrue(),
						Generic:        (&PolicyGeneric{}).GetDefaults(),
						Docker: &PolicyDocker{
							KeepSemver: utl.NewTrue(),
							Exclude: []string{
								"latest",
							},
						},
					},
					{
						Name: "policy_misc",
						Repos: []string{
							"rpm-prod-local",
							"rpm-local",
							"generic-local",
						},
						Schedule:       "*/30 * * * * *",
						Retention:      utl.NewDuration(24 * time.Hour),
						LastModified:   utl.NewTrue(),
						LastDownloaded: utl.NewTrue(),
						Generic: &PolicyGeneric{
							Include: []string{
								"prod/*",
							},
							Exclude: []string{
								"*2.2.*",
								"*2.1.0*",
							},
						},
						Docker: (&PolicyDocker{}).GetDefaults(),
					},
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			cfg, err := Load(tt.cli, Meta{})
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantData, cfg)
			if cfg != nil {
				assert.NotEmpty(t, cfg.String())
			}
		})
	}
}

func UnsetEnv(prefix string) (restore func()) {
	before := map[string]string{}

	for _, e := range os.Environ() {
		if !strings.HasPrefix(e, prefix) {
			continue
		}

		parts := strings.SplitN(e, "=", 2)
		before[parts[0]] = parts[1]

		os.Unsetenv(parts[0])
	}

	return func() {
		after := map[string]string{}

		for _, e := range os.Environ() {
			if !strings.HasPrefix(e, prefix) {
				continue
			}

			parts := strings.SplitN(e, "=", 2)
			after[parts[0]] = parts[1]

			// Check if the envar previously existed
			v, ok := before[parts[0]]
			if !ok {
				// This is a newly added envar with prefix, zap it
				os.Unsetenv(parts[0])
				continue
			}

			if parts[1] != v {
				// If the envar value has changed, set it back
				os.Setenv(parts[0], v)
			}
		}

		// Still need to check if there have been any deleted envars
		for k, v := range before {
			if _, ok := after[k]; !ok {
				// k is not present in after, so we set it.
				os.Setenv(k, v)
			}
		}
	}
}
