package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequirement(t *testing.T) {
	cases := []struct {
		req   string
		valid bool
		R     Requirement
	}{
		{"389X-Y > 1.0", true, Requirement{"389X-Y", ">", "1.0"}},
		{"a>=1.0-beta", true, Requirement{"a", ">=", "1.0-beta"}},
		{"a", true, Requirement{"a", "", ""}},
	}

	for _, c := range cases {
		require, err := NewRequirement(c.req)
		assert.Equal(t, c.valid, err == nil)
		assert.Equal(t, c.R, require)
	}
}

func TestRequirementMatch(t *testing.T) {
	cases := []struct {
		R        Requirement
		Name     string
		Version  string
		expected bool
	}{
		{Requirement{"A", ">", "1.0"}, "B", "2.0", false},
		{Requirement{"A", "", "1.0"}, "A", "2.0", true},
		{Requirement{"A", "=", ""}, "A", "2.0", true},
		{Requirement{"A", ">", "1.0"}, "A", "2.0", true},
		{Requirement{"A", ">", "1.0"}, "A", "1.1", true},
		{Requirement{"A", ">", "1.0"}, "A", "1.0.1", true},
		{Requirement{"A", ">", "1.0"}, "A", "0.9", false},
		{Requirement{"A", "<", "1.0"}, "A", "1.0.1", false},
		{Requirement{"A", "<", "1.0"}, "A", "0.9", true},
		{Requirement{"A", "<=", "1.0"}, "A", "1.0.1", false},
		{Requirement{"A", "<=", "1.0"}, "A", "1.0", true},
		{Requirement{"A", "<=", "1.0"}, "A", "0.9", true},
		{Requirement{"A", ">=", "1.0"}, "A", "1.0.1", true},
		{Requirement{"A", ">=", "1.0"}, "A", "1.0", true},
		{Requirement{"A", ">=", "1.0"}, "A", "0.9", false},
		{Requirement{"A", "=", "1.0"}, "A", "1.0", true},
		{Requirement{"A", "=", "1.0"}, "A", "1.1", false},
		// FIXME: the maintained version is a problem, it MUST be another para?
		//	{Requirement{"A", ">", "1.0"}, "A", "1.0-main", true},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, c.R.Match(c.Name, c.Version))
	}
}
