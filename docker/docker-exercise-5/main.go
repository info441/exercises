package main

import (
	"log"
	"net/http"
	"os"
)

// main is the main entrypoint for the go application.
// this function will be called when the appication starts.
func main() {

	// get the values of the required environment variables,
	// and exit if they are undefined.

	// port will be the port that the web server wil listen on
	port := getEnvOrExit("PORT")

	// STATICDIR will be the directory to serve files out of.
	// This should be an absolute path within the context of
	// your docker container.
	staticDir := getEnvOrExit("STATICDIR")

	// create a new serve mux and attach the file server route to it
	// a mux is essentially a router that will route requests to specific
	// functions (Handlers or HandlerFuncs) based on specific routes.
	mux := http.NewServeMux()

	// http.FileServer is an http.Handler that serves files out of a
	// specific directory. You don't need to know how it does that at
	// this point, but just know that using http.FileServer is essentially
	// the same as using one of your own handlers that you made.
	mux.Handle("/", http.FileServer(http.Dir(staticDir)))

	// print out that the server is listening on the given port
	// so when someone looks at the logs, we know something happened.
	log.Printf("server is listening on 0.0.0.0:%s", port)

	// start the http web server listenong from any interface on the
	// given port. "0.0.0.0" means any interface, so it will accept connections
	// from any origin on port "port". However, you do not need to specify
	// 0.0.0.0 with Go. You can leave the address blank and just provide the port.
	// ":4000" is a valid address that Go's net/http package will interpret to mean
	// "0.0.0.0:4000".
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, mux))

	// the log.Fatal wrapping the http.ListenAndServe function will report any errors
	// that were reported by http.ListenAndServe and terminate the process. This is fine
	// to terminate the process here since if our web server failed to start, we likely
	// have no reason for this application to be running anyway.
}

// getEnvOrExit gets the value of the "envName" environment variable and returns it.
// Exits the process if the given environment variable is not defined.
func getEnvOrExit(envName string) string {
	// get the environment variable "envName"
	env := os.Getenv(envName)

	// if the environment variable is undefined (length of 0)
	// then fatally log with a message noting that the environment
	// variable was not set and needs to be. Fatally logging will
	// terminate the process.
	if len(env) == 0 {
		log.Fatalf(
			"no value specified for the %s environment variable. Please set the %s envionrment variable",
			envName,
			envName)
	}
	return env
}
