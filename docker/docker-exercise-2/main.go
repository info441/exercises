package main

import (
	"log"
	"net/http"
	"os"
)

type context struct {
	port string
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("no value for env PORT found, please set a value for PORT")
	}
	http.HandleFunc("/", context{port: port}.successHandler)
	log.Println("listening on 0.0.0.0:" + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (ctx context) successHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		Success!
		
		You are able to connect to your docker container on port ` + ctx.port))
}
