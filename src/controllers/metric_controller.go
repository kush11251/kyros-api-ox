package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"kyros-api/src/models"
)

func getMetrics(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	db, err := sqlx.Connect("postgres", "user:password@localhost/database")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query the database
	var metrics []models.Metric
	err = db.Select(&metrics, "SELECT * FROM metrics")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the metrics as JSON
	json.NewEncoder(w).Encode(metrics)
}