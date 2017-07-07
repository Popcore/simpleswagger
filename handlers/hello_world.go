package handlers

import (
	"fmt"
	"net/http"
)

// swagger:route GET /pets pets users
//
// Lists pets filtered by some parameters.
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//     Security:
//       api_key:
//       oauth: read, write
//
//     Responses:
//       default: genericError
//       200: someResponse
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Swagger!")
}
