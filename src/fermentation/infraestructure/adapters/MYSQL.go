package adapters

import (
	"Integrador/src/fermentation/domain/entities"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"Integrador/src/fermentation/domain"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) domain.FermentationRepository {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(f entities.Fermentation) (entities.Fermentation, error) {
	if f.StartedAt.IsZero() {
		f.StartedAt = time.Now()
	}
	query := `INSERT INTO fermentation_records (
		operator_id, started_at, raw_material, raw_material_quantity,
		total_volume, yeast_type, sugar_concentration, mineral_content,
		temperature, agitation_rpm, ethanol_concentration, yeast_growth, ph
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := m.conn.Exec(query,
		f.OperatorID, f.StartedAt, f.RawMaterial, f.RawMaterialQuantity,
		f.TotalVolume, f.YeastType, f.SugarConcentration, f.MineralContent,
		f.Temperature, f.AgitationRPM, f.EthanolConcentration, f.YeastGrowth, f.PH,
	)

	if err != nil {
		return entities.Fermentation{}, fmt.Errorf("error saving fermentation: %v", err)
	}

	return f, nil
}

func (m *MySQL) Update(f entities.Fermentation) (entities.Fermentation, error) {
	query := `UPDATE fermentation_records SET
		raw_material = ?, raw_material_quantity = ?, total_volume = ?,
		yeast_type = ?, sugar_concentration = ?, mineral_content = ?,
		temperature = ?, agitation_rpm = ?, ethanol_concentration = ?,
		yeast_growth = ?, ph = ?
		WHERE id = ? AND operator_id = ?`

	_, err := m.conn.Exec(query,
		f.RawMaterial, f.RawMaterialQuantity, f.TotalVolume,
		f.YeastType, f.SugarConcentration, f.MineralContent,
		f.Temperature, f.AgitationRPM, f.EthanolConcentration,
		f.YeastGrowth, f.PH,
		f.ID, f.OperatorID,
	)

	if err != nil {
		return entities.Fermentation{}, fmt.Errorf("error updating fermentation: %v", err)
	}

	return f, nil
}

func (m *MySQL) GetByID(id, userID int) (entities.Fermentation, error) {
	var f entities.Fermentation

	query := `SELECT id, operator_id, started_at, raw_material, raw_material_quantity,
		total_volume, yeast_type, sugar_concentration, mineral_content,
		temperature, agitation_rpm, ethanol_concentration, yeast_growth, ph
		FROM fermentation_records WHERE id = ? AND operator_id = ?`

	err := m.conn.QueryRow(query, id, userID).Scan(
		&f.ID, &f.OperatorID, &f.StartedAt, &f.RawMaterial, &f.RawMaterialQuantity,
		&f.TotalVolume, &f.YeastType, &f.SugarConcentration, &f.MineralContent,
		&f.Temperature, &f.AgitationRPM, &f.EthanolConcentration, &f.YeastGrowth, &f.PH,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return f, errors.New("record not found")
		}
		return f, fmt.Errorf("error retrieving fermentation record: %v", err)
	}

	return f, nil
}

func (m *MySQL) GetAll(userID int) ([]entities.Fermentation, error) {
	query := `SELECT id, operator_id, started_at, raw_material, raw_material_quantity,
		total_volume, yeast_type, sugar_concentration, mineral_content,
		temperature, agitation_rpm, ethanol_concentration, yeast_growth, ph
		FROM fermentation_records WHERE operator_id = ?`

	rows, err := m.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving fermentation records: %v", err)
	}
	defer rows.Close()

	var results []entities.Fermentation

	for rows.Next() {
		var f entities.Fermentation
		err := rows.Scan(
			&f.ID, &f.OperatorID, &f.StartedAt, &f.RawMaterial, &f.RawMaterialQuantity,
			&f.TotalVolume, &f.YeastType, &f.SugarConcentration, &f.MineralContent,
			&f.Temperature, &f.AgitationRPM, &f.EthanolConcentration, &f.YeastGrowth, &f.PH,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning fermentation row: %v", err)
		}
		results = append(results, f)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating fermentation rows: %v", err)
	}

	return results, nil
}

func (m *MySQL) Delete(id, userID int) error {
	query := `DELETE FROM fermentation_records WHERE id = ? AND operator_id = ?`
	_, err := m.conn.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("error deleting fermentation record: %v", err)
	}
	return nil
}
