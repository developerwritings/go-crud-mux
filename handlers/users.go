package handlers

import (
	"encoding/json"
	"net/http"
)

type Users struct {
	Name string `json:"name,omitempty"`
}

func UserHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(nil)
}
