package handlers

import (
	"encoding/json"
	"net/http"
)

func UserHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(nil)
}
