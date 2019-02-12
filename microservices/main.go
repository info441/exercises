package main

import (
    "log"
    "net/http"
		"net/url"
		"net/http/httputil"
		"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API Gateway"))
}

func main() {
	addr := os.Getenv("ADDR")
	chatAddr := os.Getenv("CHAT")
	imageAddr := os.Getenv("IMAGE")

	if addr == "" {
		addr = ":4000"
	}

	chatProxy := httputil.NewSingleHostReverseProxy(&url.URL{ Scheme: "http", Host: chatAddr })
	imageProxy := httputil.NewSingleHostReverseProxy(&url.URL{ Scheme: "http", Host: imageAddr })

	mux := http.NewServeMux()  // create mux
	mux.HandleFunc("/", IndexHandler)
	mux.Handle("/v1/chat", chatProxy)   // register the proxies
	mux.Handle("/v1/image", imageProxy)
	log.Fatal(http.ListenAndServe(addr, mux))
}
