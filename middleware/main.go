package main

import (
	"fmt"
	"github.com/info441/exercises/middleware/middleware"
	"log"
	"net/http"
	"os"
	"time"
)


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

func UsersMeHandler(w http.ResponseWriter, r *http.Request, u *middleware.User) {
	w.Write([]byte(fmt.Sprintf("current user: %d: %s", u.ID, u.UserName)))
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = "localhost:4000"
	}
	
	
	// See how we are able to use NewAuthenticatedMux as a mutex instead of the traditional http.NewServeMux
	// This is because we composed a new type (NewAuthenticatedMux) that contains http.ServeMux
	mux := middleware.NewAuthenticatedMux()
	mux.HandleFunc("/v1/hello", HelloHandler)
	mux.HandleFunc("/v1/time", CurrentTimeHandler)
	// Again, this is a handlerFunc that we defined in authenticator.go
	// UsersMeHandler follows the function signature and therefore is allowed as a argument
	mux.HandleAuthenticatedFunc("/users/me", UsersMeHandler)
	
	// TODO: Wrap your mux in with your Logger and one of the cors middleware that you implemented.
	// Make sure to change the mux required for http.ListenAndServe
	
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
	
}