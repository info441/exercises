package middleware

import (
	"log"
	"net/http"
	"time"
)

//TODO: implement a Logger middleware handler
//that also captures and logs the response statusCode

// We'll first seek to write a middleware that allows us to log to stdout information about
// a given request. Refer to https://drstearns.github.io/tutorials/gomiddleware/

// Logger is a middleware handler
// TODO: Include a field that is of type http.Handler
type Logger struct {
	wrappedHandler http.Handler
}

// Create a "constructor" for your Logger
func NewLogger(handlerToWrap http.Handler) http.Handler {
	return &Logger{handlerToWrap}
}

// ServeHTTP handles the request by passing it to the real handler and logging the request details
// TODO: Log (log.Printf() is fine) out the request method, request url path and the duration of the request
// HINT: Use time.Since() to capture duration
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	begin := time.Now()
	l.wrappedHandler.ServeHTTP(w, r)
	log.Printf("%s %s - %d %v", r.Method, r.URL.Path, http.StatusOK, time.Since(begin))
}