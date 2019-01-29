package main

import (
	"github.com/gorilla/mux"
	"github.com/info441/exercises/postman/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0{
		addr = "localhost:4000"
	}
	
	mux := mux.NewRouter()
	context := &handlers.Context{
		Users: handlers.UsersList{},
	}
	
	mux.HandleFunc("/v1/registration", context.RegistrationHandler)
	mux.HandleFunc("/v1/login", context.LoginHandler)
	mux.HandleFunc("/v1/user/{id}", context.UsersHandler)
	//  Start a web server listening on the address you read from
	//  the environment variable, using the mux you created as
	//  the root handler. Use log.Fatal() to report any errors
	//  that occur when trying to start the web server.
	log.Printf("Server is listening at http://%s:", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
	
}



