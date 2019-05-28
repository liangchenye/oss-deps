package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/liangchenye/oss-deps/pkg"
)

// TODO: add more cases
var testPkgData = []pkg.Package{
	pkg.Package{Name: "GoodA", Version: "1.0.0",
		BuildRequires: []pkg.Requirement{{Name: "GoodB"}, {Name: "GoodC"}, {Name: "GoodD"}}},
	pkg.Package{Name: "GoodB", Version: "2.0.0",
		BuildRequires: []pkg.Requirement{{Name: "GoodD", Oper: "<", Version: "1.1"}}},
	pkg.Package{Name: "GoodC", Version: "1.1.1"},
	pkg.Package{Name: "GoodD", Version: "1.0"},
	pkg.Package{Name: "BadA", Version: "1.0.0",
		BuildRequires: []pkg.Requirement{{Name: "BadB", Oper: ">", Version: "1.0"}}},
	pkg.Package{Name: "BadB", Version: "0.9"},
}

var testTrainData = []pkg.Train{
	pkg.Train{Name: "CleanOS", Version: "1.0",
		Packages: []pkg.Package{{Name: "GoodA", Version: "1.0.0"},
			{Name: "GoodB", Version: "2.0.0"},
			{Name: "GoodC", Version: "1.1.1"},
			{Name: "GoodD", Version: "1.0"},
		}},
	pkg.Train{Name: "CleanOS", Version: "2.0",
		Packages: []pkg.Package{{Name: "BadA", Version: "1.0.0"}}},
}

func Exporter(filename string) error {
	content, _ := json.MarshalIndent(testPkgData, "\t", "  ")

	return ioutil.WriteFile(filename, content, 0644)
}

func main() {
	pkgData := "./meta"
	content, _ := json.MarshalIndent(testPkgData, "\t", "  ")
	ioutil.WriteFile(pkgData, content, 0644)

	trainData := "./trainmeta"
	content, _ = json.MarshalIndent(testTrainData, "\t", " ")
	ioutil.WriteFile(trainData, content, 0644)
}
