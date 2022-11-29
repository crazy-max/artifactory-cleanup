package config

import (
	"github.com/crazy-max/artifactory-cleanup/pkg/utl"
)

// PolicyDocker holds docker's policy configuration
type PolicyDocker struct {
	KeepSemver *bool    `yaml:"keepSemver,omitempty" json:"keepSemver,omitempty"`
	Exclude    []string `yaml:"exclude,omitempty" json:"exclude,omitempty"`
}

// GetDefaults gets the default values
func (s *PolicyDocker) GetDefaults() *PolicyDocker {
	n := &PolicyDocker{}
	n.SetDefaults()
	return n
}

// SetDefaults sets the default values
func (s *PolicyDocker) SetDefaults() {
	s.KeepSemver = utl.NewTrue()
}

// IsExcluded checks if a tag is excluded
func (s *PolicyDocker) IsExcluded(tag string) bool {
	if len(s.Exclude) == 0 {
		return false
	}
	for _, exclude := range s.Exclude {
		if utl.MatchString(exclude, tag) {
			return true
		}
	}
	return false
}
