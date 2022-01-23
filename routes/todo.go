package routes

import (
	"go-crud/handlers"
	"log"

	"github.com/gorilla/mux"
)

func TodoRouters(router *mux.Router) *mux.Router {
	log.Println("Todo Routes")
	todo := router.PathPrefix("/todo").Subrouter()
	LoadTodoSubRoutes(todo)
	return todo
}

func LoadTodoSubRoutes(router *mux.Router) *mux.Router {
	log.Println("Todo Sub Routes")
	router.HandleFunc("/", handlers.TodoHome).Methods("GET")
	return router
}
