package types

//  Success
//
// swagger:response successResponse
type SuccessResponse struct {
	// in: body
	Body ResponseBody
}

type ResponseBody struct {
	// The HTTP code
	Code int `json:"code"`

	// The response message
	Message string `json:"message"`

	// Details about the HTTP request made by the client
	Details []string `json:"details"`
}
