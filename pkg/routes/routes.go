package routes

import (
	"github.com/gorilla/mux"
	"github.com/kiyoshi-87/RetireMate-GO-backend/pkg/controllers"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	//Root Route
	r.HandleFunc("/", controllers.Home).Methods("GET")

	// Read the data
	r.HandleFunc("/data/{email}", controllers.GetUserData).Methods("GET")

	// Create the data
	r.HandleFunc("/data/", controllers.CreateUser).Methods("POST")

	// Update the data
	r.HandleFunc("/data/{email}", controllers.UpdateUserData).Methods("PUT")

	//Health Check
	r.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	return r
}
