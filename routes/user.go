package routes

import (
	"encoding/json"
	"fmt"
	"go-crud/db"
	"go-crud/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRouters(router *mux.Router) *mux.Router {
	log.Println("User Routes")
	users := router.PathPrefix("/users").Subrouter()
	LoadSubRoutes(users)
	return users
}

func LoadSubRoutes(router *mux.Router) *mux.Router {
	log.Println("User Sub Routes")
	router.HandleFunc("/", handlers.UserHome).Methods("GET")
	router.HandleFunc("/create", Create).Methods("POST")
	return router
}

func Create(w http.ResponseWriter, r *http.Request) {
	var user db.User
	fmt.Println(json.NewDecoder(r.Body).Decode(&user))
	json.NewEncoder(w).Encode(nil)
}
