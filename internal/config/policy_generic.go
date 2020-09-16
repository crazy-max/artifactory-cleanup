package config

// PolicyGeneric holds generic's policy configuration
type PolicyGeneric struct {
	Include []string `yaml:"include,omitempty" json:"include,omitempty"`
	Exclude []string `yaml:"exclude,omitempty" json:"exclude,omitempty"`
}

// GetDefaults gets the default values
func (s *PolicyGeneric) GetDefaults() *PolicyGeneric {
	n := &PolicyGeneric{}
	n.SetDefaults()
	return n
}

// SetDefaults sets the default values
func (s *PolicyGeneric) SetDefaults() {
	s.Include = []string{""}
}
