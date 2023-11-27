package handlers

import (
	"encoding/json"
	"net/http"
	"randgo/store"
)

type CompanyServiceRequest struct {
	ID uint64 `json:"id"`
}

func CompanyService(w http.ResponseWriter, r *http.Request) {
	var companyServiceRequest CompanyServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&companyServiceRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	companyService := store.DBStore.GetCompanyService(companyServiceRequest.ID)
	if companyService == nil {
		http.Error(w, "Company service not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(companyService)
}
