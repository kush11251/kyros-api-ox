package models

import (
	"time"
)

type Metric struct {
	ID       string    `json:"id"`
	Value    float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}