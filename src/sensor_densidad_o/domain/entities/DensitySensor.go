package entities

import "time"

type DensitySensor struct {
	MeasurementID int       `json:"measurement_id"`
	UserID        int       `json:"user_id"`
	Timestamp     time.Time `json:"timestamp"`
	Density       float64   `json:"density"`
}
