package api

import "github.com/gorilla/mux"

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", HealthHandler).Methods("GET")
	r.HandleFunc("/narrative", CreateNarrative).Methods("POST")
	r.HandleFunc("/narrative/{id}", GetNarrative).Methods("GET")

	return r
}
