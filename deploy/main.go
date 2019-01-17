package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const headerCORS = "Access-Control-Allow-Origin"
const corsAnyOrigin = "*"
const headerContentType = "Content-Type"
const contentTypePlainText = "text/plain"

func main() {
	ADDR := os.Getenv("ADDR")
	if len(ADDR) == 0 {
		ADDR = ":443"
	}
	
	TLSCERT := os.Getenv("TLSCERT")
	if len(TLSCERT) == 0 {
		log.Fatal("No TLSCERT environment variable found")
	}
	
	TLSKEY := os.Getenv("TLSKEY")
	if len(TLSKEY) == 0 {
		log.Fatal("No TLSKEY environment variable found")
	}
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	
	log.Printf("Server is listening at http://%s:", ADDR)
	log.Fatal(http.ListenAndServeTLS(ADDR, TLSCERT, TLSKEY, mux))
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		name = "World"
	}
	w.Header().Add(headerCORS, corsAnyOrigin)
	w.Header().Add(headerContentType, contentTypePlainText)
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}
