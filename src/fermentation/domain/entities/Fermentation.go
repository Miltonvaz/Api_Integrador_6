package entities

import "time"

type Fermentation struct {
	ID                   int       `json:"id"`
	OperatorID           int       `json:"operator_id"`
	StartedAt            time.Time `json:"started_at"`
	RawMaterial          string    `json:"raw_material"`
	RawMaterialQuantity  float64   `json:"raw_material_quantity"`
	TotalVolume          float64   `json:"total_volume"`
	YeastType            string    `json:"yeast_type"`
	SugarConcentration   float64   `json:"sugar_concentration"`
	MineralContent       string    `json:"mineral_content"`
	Temperature          float64   `json:"temperature"`
	AgitationRPM         float64   `json:"agitation_rpm"`
	EthanolConcentration float64   `json:"ethanol_concentration"`
	YeastGrowth          string    `json:"yeast_growth"`
	PH                   float64   `json:"ph"`
}
