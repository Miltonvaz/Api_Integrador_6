package entities

import "time"

type MotorFermentador struct {
	MeasurementID int       `json:"measurement_id"`
	UserID        int       `json:"user_id"`
	Timestamp     time.Time `json:"timestamp"`
	RPM           float64   `json:"rpm"`
}
