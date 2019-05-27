package pkg

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/go-version"
)

var reqRe, _ = regexp.Compile(`([\w\.\-]+)[\s]*([\=|>|>\=|<\=|<|]*)[\s]*([\w\.\-]*)`)

// Requirement: it is just a simple way
// the complicate situation will be:
// 1. (a > b) && (a<c)
// 2. (a > b) || (a<c)
// TODO: find the in CVE or other situation
// we make it simpler -- only '&&'
type Requirement struct {
	Name string
	// Oper: Now we only support  "=", ">", "<", ">=", "<="
	Oper    string
	Version string
}

func NewRequirement(req string) (require Requirement, err error) {
	strs := reqRe.FindStringSubmatch(req)
	if len(strs) == 4 {
		require.Name = strs[1]
		require.Oper = strs[2]
		require.Version = strs[3]
	} else if len(strs) == 2 {
		require.Name = strs[1]
	} else {
		err = fmt.Errorf("invalid require string '%s'", req)
	}

	return
}

func (r *Requirement) Match(Name string, Version string) bool {
	if r.Name != Name {
		return false
	}

	// if the oper or version is not set, it means there is no API specified
	if r.Oper == "" || r.Version == "" {
		return true
	}
	vr, err := version.NewVersion(r.Version)
	if err != nil {
		return false
	}

	vin, err := version.NewVersion(Version)
	if err != nil {
		return false
	}

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
