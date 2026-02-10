package main

import (
	"log"
	"net/http"
	"truesrc/api"
)

func main() {
	r := api.SetupRoutes()
	log.Println("TRUESRC running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
