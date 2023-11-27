package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"randgo/database"
	"randgo/utils"
	"strconv"

	"gorm.io/gorm"
)

type LoginRequest struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if loginRequest.Mail == "" {
		http.Error(w, "Mail is empty", http.StatusBadRequest)
		return
	}

	if loginRequest.Password == "" {
		http.Error(w, "Password is empty", http.StatusBadRequest)
		return
	}

	var userID uint64

	if err := database.Connection.Model(&utils.User{}).Where("Mail = ?", loginRequest.Mail).Select("ID").Row().Scan(&userID); err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("No rows were returned for mail: %s\n", loginRequest.Mail)
			http.Error(w, "Invalid mail or password", http.StatusUnauthorized)
			return
		} else {
			fmt.Printf("Error while getting mail: %s\n", loginRequest.Mail)
			panic(err)
		}
	}

	var sessionToken string

	var count int64
	if err := database.Connection.Model(&utils.Login{}).Where("ID = ? AND Password = ?", userID, loginRequest.Password).Count(&count).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if count == 0 {
		http.Error(w, "Invalid mail or password", http.StatusUnauthorized)
		return
	} else {
		// TODO: Generate a good session token
		sessionToken = "TOKEN" + strconv.FormatUint(userID, 2)
	}

	fmt.Printf("Logged in user ID : %s , mail: %s\n", strconv.FormatUint(userID, 10), loginRequest.Mail)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": sessionToken})
}
