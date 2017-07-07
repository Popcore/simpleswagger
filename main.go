// Package - Simple Swagger API
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
//
//     Schemes: http, https
//     Host: localhost
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: smallzed@gmail.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"github.com/popcore/simple-swagger/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HelloWorld)

	log.Fatal(http.ListenAndServe("localhost:3001", mux))
}

// A GenericError is the default error message that is generated.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}

// A GenericResponse is the default response returned.
//
// swagger:response someResponse
type GenericResponse struct {
	// in: body
	Code    int    `json:"code"`
	Message string `json:"message"`
}
