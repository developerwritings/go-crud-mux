package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-crud/db"
	"log"
	"net/http"
)

func TodoHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	log.Println("Create todo")
	var user db.User
	fmt.Println(json.NewDecoder(r.Body).Decode(&user))
	todoCollection := db.FetchTodoCollection()
	result, err := todoCollection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(result)
}
