package handlers

import (
	"encoding/json"
	"net/http"
	"randgo/store"
)

type CompanyRequest struct {
	ID uint64 `json:"id"`
}

func Company(w http.ResponseWriter, r *http.Request) {
	var companyRequest CompanyRequest
	if err := json.NewDecoder(r.Body).Decode(&companyRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if companyRequest.ID == 0 {
		http.Error(w, "ID is empty", http.StatusBadRequest)
		return
	}

	user := store.DBStore.GetCompany(companyRequest.ID)
	if user == nil {
		http.Error(w, "Company not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
