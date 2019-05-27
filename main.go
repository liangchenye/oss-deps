package main

import (
	"fmt"

	"github.com/liangchenye/oss-deps/hub"
	//	"github.com/liangchenye/oss-deps/pkg"
)

func main() {
	localDataURL := "./data/localpkgs"
	if err := hub.SetDefault("local"); err != nil {
		fmt.Println(err)
		return
	}

	h, err := hub.GetDefault()
	if err != nil {
		fmt.Println(err)
		return
	}

	h.Init(localDataURL)
	pkgs, _ := h.GetPackages()
	fmt.Println(pkgs)
}
