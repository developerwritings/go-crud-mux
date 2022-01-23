package routes

import (
	"go-crud/handlers"
	"log"

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
	return router
}
