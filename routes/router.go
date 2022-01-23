package routes

import (
	"log"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	log.Println("Creating Object for the mux package")
	router = mux.NewRouter()
}

func LoadConfig() *mux.Router {
	log.Println("Loading Router Config")
	UserRouters(router)
	TodoRouters(router)
	return router
}
