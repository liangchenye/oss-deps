package hub

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/liangchenye/oss-deps/pkg"
)

const (
	localName           = "local"
	localSourceProtocal = pkg.SourceProtocalFile
)

type LocalHub struct {
	MetaURL string
	DataDir string
}

func init() {
	Register(localName, &LocalHub{})
}

func (lh *LocalHub) Init(data map[string]string) error {
	if v, ok := data["meta-url"]; !ok {
		return errors.New("no 'meta-url' provided")
	} else {
		lh.MetaURL = v
	}

	if v, ok := data["data-dir"]; !ok {
		return errors.New("no 'data-dir' provided")
	} else {
		lh.DataDir = v
	}
	// TODO: verify if the url is valid
	return nil
}

// GetSource: this is where the standards work
func (lh *LocalHub) GetSource(p pkg.Package) (pkg.Source, error) {
	var s pkg.Source
	s.Protocal = localSourceProtocal
	// no maintainance
	// but assume that the source is well packaged as we defined ...
	if p.MVersion == "" {
		// the file name of openssl 1.1.1 SHOULD be  'openssl-1.1.1.tar'
		// it will stored at '{DATADIR}/openssl/1.1.1/'
		s.URL = filepath.Join(lh.DataDir, p.Name, p.Version, fmt.Sprintf("%s-%s.tar", p.Name, p.Version))
	} else {
		// the file name of openssl 1.1.1 maintained one time SHOULD be  'openssl-1.1.1-h1.tar'
		// it will stored at '{DATADIR}/openssl/1.1.1/'
		s.URL = filepath.Join(lh.DataDir, p.Name, p.Version, fmt.Sprintf("%s-%s-h%s.tar", p.Name, p.Version, p.MVersion))
	}

	return s, nil
}

// TODO: add a cache?
func (lh *LocalHub) GetPackages() (pkgs []pkg.Package, err error) {
	data, err := ioutil.ReadFile(lh.MetaURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &pkgs)
	if err != nil {
		return
	}

	// TODO: we can also add this information to metafile
	for i := range pkgs {
		pkgs[i].DevSource, _ = lh.GetSource(pkgs[i])
	}
	return
}
