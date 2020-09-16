package config

import (
	"encoding/json"

	"github.com/crazy-max/gonfig"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Config holds configuration details
type Config struct {
	Cli         Cli          `yaml:"-" json:"-" label:"-" file:"-"`
	Meta        Meta         `yaml:"-" json:"-" label:"-" file:"-"`
	Artifactory *Artifactory `yaml:"artifactory,omitempty" json:"artifactory,omitempty" validate:"required"`
	Policies    Policies     `yaml:"policies,omitempty" json:"policies,omitempty" validate:"required,unique=Name,dive"`
}

// Load returns Configuration struct
func Load(cli Cli, meta Meta) (*Config, error) {
	cfg := Config{
		Cli:  cli,
		Meta: meta,
	}

	fileLoader := gonfig.NewFileLoader(gonfig.FileLoaderConfig{
		Filename: cli.Cfgfile,
		Finder: gonfig.Finder{
			BasePaths: []string{
				"/etc/artifactory-cleanup/artifactory-cleanup",
				"$XDG_CONFIG_HOME/artifactory-cleanup",
				"$HOME/.config/artifactory-cleanup",
				"./artifactory-cleanup",
			},
			Extensions: []string{"yaml", "yml"},
		},
	})
	if found, err := fileLoader.Load(&cfg); err != nil {
		return nil, errors.Wrap(err, "Failed to decode configuration from file")
	} else if !found {
		log.Debug().Msg("No configuration file found")
	} else {
		log.Info().Msgf("Configuration loaded from file: %s", fileLoader.GetFilename())
	}

	envLoader := gonfig.NewEnvLoader(gonfig.EnvLoaderConfig{
		Prefix: "ATFCLNP_",
	})
	if found, err := envLoader.Load(&cfg); err != nil {
		return nil, errors.Wrap(err, "Failed to decode configuration from environment variables")
	} else if !found {
		log.Debug().Msg("No ATFCLNP_* environment variables defined")
	} else {
		log.Info().Msgf("Configuration loaded from %d environment variables", len(envLoader.GetVars()))
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) validate() error {
	return validator.New().Struct(cfg)
}

// String returns the string representation of configuration
func (cfg *Config) String() string {
	b, _ := json.MarshalIndent(cfg, "", "  ")
	return string(b)
}
