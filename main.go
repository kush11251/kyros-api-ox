package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc="/metrics", getMetrics).Methods("GET")
	r.HandleFunc("/metrics", createMetric).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}