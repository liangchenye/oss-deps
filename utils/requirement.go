package utils

import (
	"github.com/hashicorp/go-version"
)

// Requirement: it is just a simple way
// the complicate situation will be:
// 1. (a > b) && (a<c)
// 2. (a > b) || (a<c)
// TODO: find the in CVE or other situation
// we make it simpler -- only '&&'
type Requirement struct {
	Name string
	// Oper: Now we only support  "=", ">", "<"
	Oper    string
	Version string
}

func (r *Requirement) Match(Name string, Version string) bool {
	if r.Name != Name {
		return false
	}

	vr, _ := version.NewVersion(r.Version)
	vin, _ := version.NewVersion(Version)

	if r.Oper == "=" && vin.Equal(vr) {
		return true
	} else if r.Oper == ">" && vin.GreaterThan(vr) {
		return true
	} else if r.Oper == ">=" && vin.GreaterThanOrEqual(vr) {
		return true
	} else if r.Oper == "<" && vin.LessThan(vr) {
		return true
	} else if r.Oper == "<=" && vin.LessThanOrEqual(vr) {
		return true
	}

	return false
}
