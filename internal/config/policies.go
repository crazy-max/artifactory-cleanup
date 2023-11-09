package config

import (
	"time"

	"github.com/crazy-max/artifactory-cleanup/pkg/utl"
)

// Policies holds slice of policy
type Policies []Policy

// Policy holds policy configuration
type Policy struct {
	Name           string         `yaml:"name,omitempty" json:"name,omitempty" validate:"required"`
	Repos          []string       `yaml:"repos,omitempty" json:"repos,omitempty" validate:"required"`
	Schedule       string         `yaml:"schedule,omitempty" json:"schedule,omitempty" validate:"required"`
	Retention      *time.Duration `yaml:"retention,omitempty" json:"retention,omitempty" validate:"required"`
	RetentionCount *int           `yaml:"retentionCount,omitempty" json:"retentionCount,omitempty" validate:"required"`
	LastModified   *bool          `yaml:"lastModified,omitempty" json:"lastModified,omitempty" validate:"required"`
	LastDownloaded *bool          `yaml:"lastDownloaded,omitempty" json:"lastDownloaded,omitempty" validate:"required"`
	Common         *PolicyCommon  `yaml:"common,omitempty" json:"common,omitempty"`
	Docker         *PolicyDocker  `yaml:"docker,omitempty" json:"docker,omitempty"`
}

// GetDefaults gets the default values
func (s *Policy) GetDefaults() *Policy {
	n := &Policy{}
	n.SetDefaults()
	return n
}

// SetDefaults sets the default values
func (s *Policy) SetDefaults() {
	s.LastModified = utl.NewTrue()
	s.LastDownloaded = utl.NewTrue()
	s.Common = (&PolicyCommon{}).GetDefaults()
	s.Docker = (&PolicyDocker{}).GetDefaults()
}
