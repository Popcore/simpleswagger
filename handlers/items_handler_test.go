package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItemsHandler(t *testing.T) {

	testcases := []struct {
		query        string
		httpStatus   int
		responseBody string
	}{
		{
			query:        "",
			httpStatus:   http.StatusOK,
			responseBody: `{"items":[{"id":1,"name":"item one","colour":"red"},{"id":2,"name":"item two","colour":"yellow"}]}`,
		},
		{
			query:        "?q=1",
			httpStatus:   http.StatusOK,
			responseBody: `{"items":[{"id":1,"name":"item one","colour":"red"}]}`,
		},
		{
			query:        "?q=123",
			httpStatus:   http.StatusNotFound,
			responseBody: `{"error":{"code":404,"message":"The requested item was not found"}}`,
		},
	}

	for _, testcase := range testcases {

		queryURL := fmt.Sprintf("/items%s", testcase.query)

		req := httptest.NewRequest("GET", queryURL, nil)

		recorder := httptest.NewRecorder()

		ItemsHander(recorder, req)

		if recorder.Code != testcase.httpStatus {
			t.Errorf("Unexpected http status code: expected %d, got %d", testcase.httpStatus, recorder.Code)
		}

		if recorder.Body.String() != testcase.responseBody {
			t.Errorf("Unexpected response body: expected %s, got %s", testcase.responseBody, recorder.Body.String())
		}
	}

}
