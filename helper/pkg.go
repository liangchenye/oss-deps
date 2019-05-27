package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/liangchenye/oss-deps/pkg"
)

var testPkgData = []pkg.Package{
	pkg.Package{Name: "A", Version: "1.0.0",
		BuildRequires: []pkg.Requirement{{Name: "B"}}},
	pkg.Package{Name: "B", Version: "2.0.0"},
}

func Exporter(filename string) error {
	content, _ := json.MarshalIndent(testPkgData, "\t", "  ")

	return ioutil.WriteFile(filename, content, 0644)
}

func main() {
	output := "/tmp/localpkgs"
	if err := Exporter(output); err != nil {
		fmt.Println(err)
	}
}
