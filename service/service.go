package service

import (
	"fmt"

	"github.com/liangchenye/oss-deps/pkg"
)

type Service interface {
	Init(config map[string]string) error
	GetTrains() ([]pkg.Train, error)
	GetPackagesFromTrain(train pkg.Train) ([]pkg.Package, error)
	// Get all packages
	GetPackages() ([]pkg.Package, error)
}

var (
	defaultName string
	services    = make(map[string]Service)
)

// TODO: add a mutex
func Register(name string, h Service) error {
	services[name] = h
	return nil
}

// SetDefault
func SetDefault(name string) error {
	for storedName := range services {
		if storedName == name {
			defaultName = name
			return nil
		}
	}

	return fmt.Errorf("cannot set '%s' to default, this service is not registed", defaultName)
}

// Default
func GetDefault() (Service, error) {
	for storedName, h := range services {
		if defaultName == storedName {
			return h, nil
		}
	}
	var h Service
	return h, fmt.Errorf("cannot get the default service service")
}
