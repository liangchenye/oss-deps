package main

import (
	"encoding/json"
	"fmt"

	"github.com/liangchenye/oss-deps/hub"
	"github.com/liangchenye/oss-deps/pkg"
)

func main() {
	localConfig := make(map[string]string)
	localConfig["meta-url"] = "./data/local/meta"
	localConfig["data-dir"] = "./data/local"
	if err := hub.SetDefault("local"); err != nil {
		fmt.Println(err)
		return
	}

	h, err := hub.GetDefault()
	if err != nil {
		fmt.Println(err)
		return
	}

	h.Init(localConfig)
	pkgs, _ := h.GetPackages()
	var goodPkg pkg.Package
	var badPkg pkg.Package

	for _, p := range pkgs {
		if p.Name == "GoodA" {
			goodPkg = p
		} else if p.Name == "BadA" {
			badPkg = p
		}
	}

	errs := goodPkg.BRTree(pkgs)

	goodBytes, _ := json.MarshalIndent(goodPkg, "\t", " ")
	fmt.Println(string(goodBytes))

	if errs == nil {
		goodPkg.PrettyDebug(0)

		fmt.Println("turn into list")
		goodList := goodPkg.BRList()
		for _, p := range goodList {
			fmt.Println(p.Name, p.Version)
		}
	} else {
		fmt.Println(goodPkg)
	}

	fmt.Println("\n\n")
	errs = badPkg.BRTree(pkgs)
	if errs == nil {
		fmt.Println(errs)
	} else {
		badPkg.PrettyDebug(0)
	}
}
