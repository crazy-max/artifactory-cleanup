package config

// PolicyCommon holds common's policy configuration
type PolicyCommon struct {
	Include []string `yaml:"include,omitempty" json:"include,omitempty"`
	Exclude []string `yaml:"exclude,omitempty" json:"exclude,omitempty"`
}

// GetDefaults gets the default values
func (s *PolicyCommon) GetDefaults() *PolicyCommon {
	n := &PolicyCommon{}
	n.SetDefaults()
	return n
}

// SetDefaults sets the default values
func (s *PolicyCommon) SetDefaults() {
	s.Include = []string{""}
}
