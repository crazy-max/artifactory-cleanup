package config

import (
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
					APIKey:  "01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ",
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
						Common:         (&PolicyCommon{}).GetDefaults(),
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
						Common: &PolicyCommon{
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
