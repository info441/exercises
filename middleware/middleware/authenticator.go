package middleware

import "net/http"

// User represents a user of the system.
type User struct {
	ID       int
	UserName string
}

// GetAuthenticatedUser is a function that returns the
// current user given a request, or nil if the user is
// not currently authenticated. This is just for demo
// purposes: normally you would use your sessions package
// to get the currently authenticated user.
func GetAuthenticatedUser(r *http.Request) (*User, error) {
	return &User{1, "test"}, nil
}

// A type for authenticated handler functions hat take a `*User` as a third parameter.
type AuthenticatedHandlerFunc func(w http.ResponseWriter, r *http.Request, u *User)

// Adding this extra parameter means that this function no longer conforms to the HTTP handler function signature,
// so to use this with http.ServeMux, we need to adapt it.

// This is an adapter function that can adapt an authenticated handler function into a regular http handler function.
// TODO: Make sure to include your a call to GetAuthenticatedUser as well as AuthenticatedHandlerFunc (since it's a func)
func EnsureAuthentication(handlerFunc AuthenticatedHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := GetAuthenticatedUser(r)
		if err != nil {
			http.Error(w, "please sign in", http.StatusUnauthorized)
			return
		}
		handlerFunc(w, r, user)
	}
}

// Struct that allows this type to inherit all of the methods of http.ServeMux
type AuthenticatedMux struct {
	http.ServeMux
}

// Constructor for a new instance of AuthenticatedMux
func NewAuthenticatedMux() *AuthenticatedMux {
	return &AuthenticatedMux{}
}

// Newly defined HandlerFunc that simulates that of http.HandleFunc
func (am *AuthenticatedMux) HandleAuthenticatedFunc(pattern string, handlerFunc AuthenticatedHandlerFunc) {
	am.HandleFunc(pattern, EnsureAuthentication(handlerFunc))
}
