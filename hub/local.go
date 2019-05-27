package hub

import (
	"encoding/json"
	"io/ioutil"

	"github.com/liangchenye/oss-deps/pkg"
)

const (
	localName = "local"
)

type LocalHub struct {
	DataURL string
}

func init() {
	Register(localName, &LocalHub{})
}

func (lh *LocalHub) Init(url string) error {
	lh.DataURL = url

	// TODO: verify if the url is valid
	return nil
}

// TODO: add a cache?
func (lh *LocalHub) GetPackages() (pkgs []pkg.Package, err error) {
	data, err := ioutil.ReadFile(lh.DataURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &pkgs)
	if err != nil {
		return
	}

	return
}
