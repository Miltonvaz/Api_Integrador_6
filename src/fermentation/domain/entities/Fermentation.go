package entities

import "time"

type Fermentation struct {
	ID                     int       `json:"id"`
	OperatorID             int       `json:"operator_id"`
	StartedAt              time.Time `json:"started_at"`
	DurationHours          float64   `json:"duration_hours"`
	RawMaterial            string    `json:"raw_material"`
	SugarConcentration     float64   `json:"sugar_concentration"`
	InitialVolume          float64   `json:"initial_volume"`
	MicroorganismCategory  string    `json:"microorganism_category"`
	MicroorganismName      string    `json:"microorganism_name"`
	MicroorganismQuantity  float64   `json:"microorganism_quantity"`
	AgitationRPM           float64   `json:"agitation_rpm"`
	Temperature            float64   `json:"temperature"`
	InitialPH              float64   `json:"initial_ph"`
	FinalPH                float64   `json:"final_ph"`
	EthanolConcentration   float64   `json:"ethanol_concentration"`
	FermentationEfficiency float64   `json:"fermentation_efficiency"`
	FermentationRate       float64   `json:"fermentation_rate"`
}
