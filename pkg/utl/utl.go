package utl

import (
	"os"
	"regexp"
	"time"
)

// MatchString reports whether a string s
// contains any match of a regular expression.
func MatchString(exp string, s string) bool {
	re, err := regexp.Compile(exp)
	if err != nil {
		return false
	}
	return re.MatchString(s)
}

// GetSecret retrieves secret's value from plaintext or filename if defined
func GetSecret(plaintext, filename string) (string, error) {
	if plaintext != "" {
		return plaintext, nil
	} else if filename != "" {
		b, err := os.ReadFile(filename)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return "", nil
}

// NewFalse returns a false bool pointer
func NewFalse() *bool {
	b := false
	return &b
}

// NewTrue returns a true bool pointer
func NewTrue() *bool {
	b := true
	return &b
}

// NewDuration returns a duration pointer
func NewDuration(duration time.Duration) *time.Duration {
	return &duration
}
