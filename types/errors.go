package types

import "net/http"

// ErrorResponse is the object used to wrap errors information
type ErrorResponse struct {
	// The HTTP status code
	HTTPStatus int `json:"code"`

	// The error message
	// required: true
	Message string `json:"message"`
}

// Error Not Found
//
// swagger:response errorNotFound
type ErrorNotFound struct {
	// in: body
	Body Err
}

type Err struct {
	E ErrorResponse `json:"error"`
}

// NewErrorNotFound returns a ErrorNotFound struct populate with the
// message argument.
func NewErrorNotFound(message string) ErrorNotFound {
	return ErrorNotFound{
		Body: Err{
			E: ErrorResponse{
				HTTPStatus: http.StatusNotFound,
				Message:    message,
			},
		},
	}
}
