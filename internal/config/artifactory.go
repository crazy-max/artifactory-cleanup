package config

import (
	"time"

	"github.com/crazy-max/artifactory-cleanup/pkg/utl"
)

// Artifactory holds Artifactory details
type Artifactory struct {
	URL        string         `yaml:"url,omitempty" json:"url,omitempty" validate:"required"`
	APIKey     string         `yaml:"apiKey,omitempty" json:"apiKey,omitempty" validate:"required"`
	APIKeyFile string         `yaml:"apiKeyFile,omitempty" json:"apiKeyFile,omitempty" validate:"omitempty,file"`
	Timeout    *time.Duration `yaml:"timeout,omitempty" json:"timeout,omitempty" validate:"required"`
}

// GetDefaults gets the default values
func (s *Artifactory) GetDefaults() *Artifactory {
	n := &Artifactory{}
	n.SetDefaults()
	return n
}

// SetDefaults sets the default values
func (s *Artifactory) SetDefaults() {
	s.Timeout = utl.NewDuration(10 * time.Second)
}
