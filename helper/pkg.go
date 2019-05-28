package main

import (
	"encoding/json"
	"fmt"
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

func Exporter(filename string) error {
	content, _ := json.MarshalIndent(testPkgData, "\t", "  ")

	return ioutil.WriteFile(filename, content, 0644)
}

func main() {
	output := "./meta"
	if err := Exporter(output); err != nil {
		fmt.Println(err)
	}
}
