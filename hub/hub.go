package hub

import (
	"fmt"

	"github.com/liangchenye/oss-deps/pkg"
)

type Hub interface {
	Init(url string) error
	GetPackages() ([]pkg.Package, error)
}

var (
	defaultName string
	hubs        = make(map[string]Hub)
)

// TODO: add a mutex
func Register(name string, h Hub) error {
	hubs[name] = h
	return nil
}

// SetDefault
func SetDefault(name string) error {
	for storedName := range hubs {
		if storedName == name {
			defaultName = name
			return nil
		}
	}

	return fmt.Errorf("cannot set '%s' to default, this service is not registed", defaultName)
}

// Default
func GetDefault() (Hub, error) {
	for storedName, h := range hubs {
		if defaultName == storedName {
			return h, nil
		}
	}
	var h Hub
	return h, fmt.Errorf("cannot get the default hub service")
}
