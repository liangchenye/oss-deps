package pkg

import (
	"fmt"
)

const (
	SourceProtocalFile = "file"
	SourceProtocalGit  = "git"
)

type Source struct {
	Protocal string
	URL      string
}

// Package is used for orchestration
// it is just like rpm spec
// whether a package is ready to build or with correct source codes is not defined here
type Package struct {
	// In a repo/train, the combination of 'name-version-mversion' is unique
	// The name of a software, same with the upstream
	Name string
	// The version of a software, same with the upstream.
	// version may or may not follows 'major-minor-patch',
	//   we may not know the 'API-compatibility' from this
	// TODO plugin?
	Version string
	// MVersion: maintainance version
	// the rule of maintaince is the API/ABI must be compatible, so we can get two things:
	// 1. the higher the better
	// 2. the API must be compatible
	MVersion string

	// BuildRequires: the required package name/version/... in building
	BuildRequires []Requirement
	// Requires: the required package name/version... in running
	Requires []Requirement

	brLeaf []Package
	rLeaf  []Package

	DevSource Source
}

// BRtree: generate the tree of build requires
func (p *Package) BRTree(pkgData []Package) []error {
	// TODO: using multip error
	var errs []error
	for _, br := range p.BuildRequires {
		subPkg, err := br.Find(pkgData)
		if err != nil {
			errs = append(errs, err)
		} else {
			subErrs := subPkg.BRTree(pkgData)
			if subErrs != nil {
				errs = append(errs, subErrs...)
			} else {
				p.brLeaf = append(p.brLeaf, subPkg)
			}
		}
	}
	return errs
}

func (p *Package) PrettyDebug(tabs int) {
	for i := 0; i < tabs; i++ {
		fmt.Printf("\t")
	}
	fmt.Println(p.Name, p.Version)
	for _, br := range p.brLeaf {
		br.PrettyDebug(tabs + 1)
	}
}
