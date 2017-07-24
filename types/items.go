package types

// Item represents the item for this application
//
// An item is a generic data resource used by this demo application.
//
// An item must have an id and a name, as well as an optional colour property.
//
// swagger:model
type Item struct {
	// the item identifier
	// required: true
	// min: 1
	ID int `json:"id"`

	// the name of the item
	// required: true
	// min length: 3
	Name string `json:"name"`

	// the colour of the item
	Colour string `json:"colour"`
}

//  Success
//
// swagger:response itemResponse
type ItemResponse struct {
	// in: body
	Body Data
}

type Data struct {
	Items []Item `json:"items"`
}
