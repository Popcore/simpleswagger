package handlers

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/thingful/simpleswagger/types"
)

// ItemsHander is responsible for retuning a list of items, optionally filtered
// by ID.
//
// swagger:route GET /items items items-operation
//
// Returns a collection of items item description.
// The items returned can be optionally filtered by their identifier.
//
// Responses:
// 200: itemResponse
// 404: errorNotFound
func ItemsHander(w http.ResponseWriter, r *http.Request) {
	var items []types.Item

	reqParam := ItemRequestParam{}
	reqParam.Query = r.URL.Query().Get("q")

	itemsDB := types.NewItemDatabase()

	if reqParam.Query != "" {

		item, ok := itemsDB[reqParam.Query]

		if !ok {

			errResponse := types.NewErrorNotFound("The requested item was not found")
			bytes, err := json.Marshal(errResponse.Body)
			if err != nil {
				panic("Unexpected error marshalling response")
			}

			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			w.Write(bytes)
			return
		}

		items = append(items, item)

	} else {

		var keys []string
		for k := range itemsDB {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, key := range keys {
			items = append(items, itemsDB[key])
		}
	}

	itemResponse := types.ItemResponse{
		Body: types.Data{
			Items: items,
		},
	}

	bytes, err := json.Marshal(itemResponse.Body)
	if err != nil {
		panic("Unexpected error marshalling response")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	return
}

// ItemRequestParams is the type which use to capture the request path
//
// swagger:parameters items-operation
type ItemRequestParam struct {
	// the item unique identifier
	//
	// required: true
	// in: query
	Query string `json:"q"`
}
