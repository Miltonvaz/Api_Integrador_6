package entities

import "time"

type TurbiditySensor struct {
	MeasurementID int       `json:"measurement_id"`
	UserID        int       `json:"user_id"`
	Timestamp     time.Time `json:"timestamp"`
	Turbidity     float64   `json:"turbidity"`
}
