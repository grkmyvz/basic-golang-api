package api

import (
	"fmt"
	"log"
	"net/http"

	"randgo/api/handlers"
	"randgo/store"

	"github.com/gorilla/mux"
)

func APIServer() {

	store.DBStore = store.NewStore()

	go store.ClearStore()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.Home).Methods("GET")
	router.HandleFunc("/user", handlers.User).Methods("POST")
	router.HandleFunc("/company", handlers.Company).Methods("POST")
	router.HandleFunc("/company-service", handlers.CompanyService).Methods("POST")
	router.HandleFunc("/appointment", handlers.Appointment).Methods("POST")
	router.HandleFunc("/comment", handlers.Comment).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// TODO: Added PUT and DELETE methods
	// TODO: Add authentication middleware
	// TODO: Origin policy

	fmt.Println("Server is running on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router))

}
