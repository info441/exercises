package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIdenticonHandler(t *testing.T) {
	//This function uses the methods and structs in the
	//net/http/httptest package to build the
	//http.ResponseWriter and *http.Request that
	//will invoke the handler function,
	//and examine what was written to the response.
	
	// Add more test cases for other Headers and Status Code
	
	cases := []struct {
		name                string
		query               string
		expectedStatusCode  int
		expectedContentType string
	}{
		{
			"Valid Name Param",
			"name=test",
			http.StatusOK,
			contentTypePNG,
		},
	}
	
	for _, c := range cases {
		URL := fmt.Sprintf("/identicon?%s", c.query)
		req := httptest.NewRequest("GET", URL, nil)
		respRec := httptest.NewRecorder()
		IdenticonHandler(respRec, req)
		
		resp := respRec.Result()
		//check the response status code
		if resp.StatusCode != c.expectedStatusCode {
			t.Errorf("case %s: incorrect status code: expected %d but got %d",
				c.name, c.expectedStatusCode, resp.StatusCode)
		}
		
	}
	
}