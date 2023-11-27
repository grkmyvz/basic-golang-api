package handlers

import (
	"encoding/json"
	"net/http"
	"randgo/store"
)

type CommentRequest struct {
	ID uint64 `json:"id"`
}

func Comment(w http.ResponseWriter, r *http.Request) {
	var commentRequest CommentRequest
	if err := json.NewDecoder(r.Body).Decode(&commentRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment := store.DBStore.GetComment(commentRequest.ID)
	if comment == nil {
		http.Error(w, "Comment not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}
