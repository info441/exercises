package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"github.com/info441/exercises/zip/zipserver/models"
)

//ZipIndexHandler handles requests that should
//return a slice of *Zip records for a given key
type ZipIndexHandler struct {
	index models.ZipIndex
}

//NewZipIndexHandler constructs a new ZipIndexHandler.
//The `index` parameter will be used to get the zips
//for the requested key.
func NewZipIndexHandler(index models.ZipIndex) *ZipIndexHandler {
	return &ZipIndexHandler{
		index: index,
	}
}

//ServeHTTP handles the HTTP requests
func (zih *ZipIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//resource path will be like: /zips/city/seattle
	//get the last part of the path and convert it
	//to lower-case
	key := path.Base(r.URL.Path)
	key = strings.ToLower(key)
	
	//tell the client that the response body is JSON
	//and allow cross-origin requests
	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(headerAccessControlAllowOrigin, "*")
	
	//get the zips for the requested key
	//and encode them into JSON
	zips := zih.index[key]
	json.NewEncoder(w).Encode(zips)
}

//RootHandler handles requests for the root resource
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(headerAccessControlAllowOrigin, "*")
	fmt.Fprintf(w, "Try requesting /zips/city/seattle")
}
