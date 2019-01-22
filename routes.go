package main

import (
	"net/http"

	"github.com/prakashpandey/sample-go-webapp/auth"
	"github.com/prakashpandey/sample-go-webapp/db"
	"github.com/prakashpandey/sample-go-webapp/index"
	"github.com/prakashpandey/sample-go-webapp/user"
)

// database will be almost same for every handler
var database = db.DB{}

// Add all handlers here
var userHandler = user.UserHandler{
	Dao: database,
}

func (s *Server) routes() {
	// define all routes here
	http.HandleFunc("/", index.HelloHandler)
	http.HandleFunc("/user", s.mustAuth(user.CreateUserHandler))
	http.HandleFunc("/user/delete", s.mustAuth(userHandler.DeleteUserHandler))
}

// authenticate all protected routes
func (s *Server) mustAuth(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth.Validate("elon@spacex.com", "WayToMars") {
			fn(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Un-Authorized user"))
		return
	})
}
