package handlers

import (
	"encoding/json"
	"net/http"
	"randgo/store"
)

type AppointmentRequest struct {
	ID uint64 `json:"id"`
}

func Appointment(w http.ResponseWriter, r *http.Request) {
	var appointmentRequest AppointmentRequest
	if err := json.NewDecoder(r.Body).Decode(&appointmentRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	appointment := store.DBStore.GetAppointment(appointmentRequest.ID)
	if appointment == nil {
		http.Error(w, "Appointment not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(appointment)
}
