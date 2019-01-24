package handlers

import (
	"fmt"
	"image/png"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIdenticonHandler(t *testing.T) {
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
		{
			"Empty Query String",
			"",
			http.StatusBadRequest,
			"text/plain; charset=utf-8",
		},
		{
			"Name Param All Spaces",
			"name=%20%20",
			http.StatusBadRequest,
			"text/plain; charset=utf-8",
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
		
		//check allow origin header
		allowedOrigin := resp.Header.Get(headerAccessControlAllowOrigin)
		if allowedOrigin != originAny {
			t.Errorf("case %s: incorrect CORS header: expected %s but got %s",
				c.name, originAny, allowedOrigin)
		}
		
		//check the content type header
		contentType := resp.Header.Get(headerContentType)
		if contentType != c.expectedContentType {
			t.Errorf("case %s: incorrect Content-Type header: expected %s but got %s",
				c.name, c.expectedContentType, contentType)
		}
		
		if resp.StatusCode == http.StatusOK {
			if _, err := png.Decode(resp.Body); err != nil {
				t.Errorf("case %s: error decoding response PNG: %v", c.name, err)
			}
		}
	}
}