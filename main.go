package main

import (
	"encoding/json"
	"fmt"

	"github.com/liangchenye/oss-deps/pkg"
	"github.com/liangchenye/oss-deps/service"
)

func main() {
	localConfig := make(map[string]string)
	localConfig["train-meta-url"] = "./test/local/trainmeta"
	localConfig["meta-url"] = "./test/local/meta"
	localConfig["data-dir"] = "./test/local"
	if err := service.SetDefault("local"); err != nil {
		fmt.Println(err)
		return
	}

	h, err := service.GetDefault()
	if err != nil {
		fmt.Println(err)
		return
	}

	h.Init(localConfig)

	trains, _ := h.GetTrains()
	pkgs, _ := h.GetPackagesFromTrain(trains[0])
	fmt.Println(pkgs)

	fmt.Println("---------------")

	pkgs, _ = h.GetPackages()
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
			fmt.Println(p.Name, p.Version, p.DevSource)
		}
	} else {
		fmt.Println(goodPkg, goodPkg.DevSource)
	}

	fmt.Println("\n\n")
	errs = badPkg.BRTree(pkgs)
	if errs == nil {
		fmt.Println(errs)
	} else {
		badPkg.PrettyDebug(0)
	}
}
