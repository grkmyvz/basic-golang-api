package handlers

import (
	"encoding/json"
	"net/http"
	"randgo/store"
)

type UserRequest struct {
	ID uint64 `json:"id"`
}

func User(w http.ResponseWriter, r *http.Request) {
	var userRequest UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := store.DBStore.GetUser(userRequest.ID)
	if user == nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
