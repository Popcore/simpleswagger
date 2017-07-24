// Package - Simple Swagger API
//
// This application is a simple HTTP server documented using Swagger.
//
// It will demonstrate some of possible comment annotations and features
// that are available to turn go code into a fully compliant swagger 2.0 spec.
//
//
// Schemes: http, https
// Host: localhost
// Version: 0.0.1
// License: MIT http://opensource.org/licenses/MIT
// Contact: marco@thingful.net
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"github.com/thingful/simpleswagger/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/items", handlers.ItemsHander)

	log.Fatal(http.ListenAndServe("localhost:3001", mux))
}
