package entities

import "time"

type TemperatureSensor struct {
	MeasurementID int       `json:"measurement_id"`
	UserID        int       `json:"user_id"`
	Timestamp     time.Time `json:"timestamp"`
	Temperature   float64   `json:"temperature"`
}
