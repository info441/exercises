package main

import "net/http"

import "log"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./html")))
	log.Println("listening on 0.0.0.0:4000...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
