package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-crud/db"
	"log"
	"net/http"
)

type Users struct {
	Name string `json:"name,omitempty"`
}

func UserHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Create User")
	var user db.User
	fmt.Println(json.NewDecoder(r.Body).Decode(&user))
	userCollection := db.FetchUsersCollection()
	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(result)
}
