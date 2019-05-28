package main

import (
	"gopkg.in/macaron.v1"

	h "github.com/liangchenye/oss-deps/cmd/server/handler"
)

// SetRouters
func SetRouters(m *macaron.Macaron) {
	// Web API
	//	m.Get("/", h.IndexMetaV1Handler)

	// Trains Discovery
	m.Group("/train", func() {
		// List files
		m.Get("/", h.TrainListHandler)
		// Create a train
		m.Post("/:name/:version", h.TrainCreateHandler)
		// Get package data of a train
		m.Get("/:name/:version", h.TrainPackageHandler)
	})

}
