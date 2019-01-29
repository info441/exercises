package handlers

/*
	WARNING: Code is intentionally missing important conditional checks for request headers as well as
	incorrect session management. Please use this as a REFERENCE, don't assume that the code is
	completely fault tolerant.
 */
import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

const AuthorizationHeader = "Authorization"
const ContentTypeHeader = "Content-Type"
const ContentTypeApplicationJSON = "application/json"

type RegistrationCredentials struct {
	Email     	string 	`json:"email"`
	Password 		string 	`json:"password"`
	FirstName 	string 	`json:"firstName"`
	LastName 		string 	`json:"lastName"`
	Description string 	`json:"description"`
}


type Context struct {
	Users UsersList `json:"users"`
}

// Acts as an in memory data structure to keep track of users
// Normally you want to substitute this for data store such as MYSQL, Mongo or Redis,
// depending on the purpose
type UsersList []*User

type User struct {
	ID 					int64  `json:"id"`
	Email     	string 	`json:"-"`
	Password 		string 	`json:"-"`
	FirstName 	string 	`json:"firstName"`
	LastName 		string 	`json:"lastName"`
	Description string 	`json:"description"`
}

func (ctx *Context) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		
		// Here I'm decoding the JSON given in the request body into a RegistrationCredentials struct
		// If the client doesn't pass in json formatted body then we will respond with a bad request error
		rc := &RegistrationCredentials{}
		if err := json.NewDecoder(r.Body).Decode(rc); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err),
				http.StatusBadRequest)
			return
		}
		
		// Does linear search on whether email has been already used (this type of search is very inefficient)
		for _, u := range ctx.Users {
			if u.Email == rc.Email {
				http.Error(w, fmt.Sprintf("Account already created with given email: %s", rc.Email),
					http.StatusBadRequest)
				return
			}
		}
		
		// Using the length of the slice of users as a means of ID for each user
		id := int64(len(ctx.Users) + 1)
		user := &User{
			ID:          id,
			Email: 			 rc.Email,
			Password:    rc.Password,
			FirstName:   rc.FirstName,
			LastName:    rc.LastName,
			Description: rc.Description,
		}
		
		// Addes the current user to our in memory data store
		ctx.Users = append(ctx.Users, user)
		
		// Sets appropriate headers prior to encoding JSON to the client
		w.Header().Add(ContentTypeHeader, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
		fmt.Printf("A user was just created: email: %s\n", user.Email)
	default:
		// Handles any methods not allowed at this resource path
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}

// Strucut that represents the required JSON passed in from the client
type LoginCredentials struct {
	ID 					int64  `json:"id"`
	Email     	string 	`json:"email"`
	Password 		string 	`json:"password"`
}

func (ctx *Context) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Here I'm decoding the JSON given in the request body into a LoginCredentials struct
		// If the user doesn't pass in json formatted body then we will respond with a bad request error
		lc := &LoginCredentials{}
		if err := json.NewDecoder(r.Body).Decode(lc); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err),
				http.StatusBadRequest)
			return
		}
		
		// Grabs the user from the in memory data store (slice) and checks whether the provided
		// email and passwords match up
		user := ctx.Users[lc.ID - 1]
		if user.Email != lc.Email {
			http.Error(w, fmt.Sprintf("Invalid email credentials for given id"),
				http.StatusBadRequest)
			return
		}
		
		if user.Password != lc.Password {
			http.Error(w, fmt.Sprintf("Invalid password credentials for given id"),
				http.StatusBadRequest)
			return
		}
		
		w.Header().Add(AuthorizationHeader, "Bearer postmanIsCool")
		w.Header().Add(ContentTypeHeader, ContentTypeApplicationJSON)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
		fmt.Printf("A user just logged in: email: %s\n", user.Email)
	default:
		// Handles any methods not allowed at this resource path
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}

// JSON required from the user in order to update account first and last name
type UpdateCredentials struct {
	FirstName 	string 	`json:"firstName"`
	LastName 		string 	`json:"lastName"`
}

func (ctx *Context) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Grabs the Authorization header value and sees if correct credentials are provided
		authHeader := r.Header.Get(AuthorizationHeader)
		if len(authHeader) == 0 {
			http.Error(w, fmt.Sprintf("Header Authorization required for this resource path"),
				http.StatusBadRequest)
			return
		}
		
		// Remember that this is essentially your session token: Bearer <session token>
		if authHeader != "Bearer postmanIsCool" {
			http.Error(w, fmt.Sprintf("Header Authorization value provided is incorrect"),
				http.StatusBadRequest)
			return
		}
		
		// Grabs the id from the resource path
		idString := path.Base(r.URL.Path)
		idIndex, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Did not provide {id} as a number in /v1/user/{id}, please provide the correct ID"),
				http.StatusBadRequest)
			return
		}
		
		if idIndex - 1 >= len(ctx.Users) {
			http.Error(w, fmt.Sprintf("ID provided does not exist in data store"),
				http.StatusBadRequest)
			return
		}
		// Grabs and returns the user information to the client
		user := ctx.Users[idIndex - 1]
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
		fmt.Printf("User with ID: %d requested account information \n", user.ID)
	case http.MethodPatch:
		// Redundant code similar to up top. Should be refactored for production purposes
		authHeader := r.Header.Get(AuthorizationHeader)
		if len(authHeader) == 0 {
			http.Error(w, fmt.Sprintf("Header Authorization required for this resource path"),
				http.StatusBadRequest)
			return
		}
		
		if authHeader != "Bearer postmanIsCool" {
			http.Error(w, fmt.Sprintf("Header Authorization value provided is incorrect"),
				http.StatusBadRequest)
			return
		}
		
		idString := path.Base(r.URL.Path)
		idIndex, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Did not provide {id} as a number in /v1/user/{id}, please provide the correct ID"),
				http.StatusBadRequest)
			return
		}
		
		if idIndex - 1 >= len(ctx.Users) {
			http.Error(w, fmt.Sprintf("ID provided does not exist in data store"),
				http.StatusBadRequest)
			return
		}
		
		user := ctx.Users[idIndex - 1]
		
		uc := &UpdateCredentials{}
		if err := json.NewDecoder(r.Body).Decode(uc); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err),
				http.StatusBadRequest)
			return
		}
		
		// Updates the fields in the User struct
		user.FirstName = uc.FirstName
		user.LastName = uc.LastName
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
		fmt.Printf("A user recently updated their name: %s %s", user.FirstName, user.LastName)
	default:
		// Handles any methods not allowed at this resource path
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}