package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"Integrador/src/sensor_turbuidez/domain"
	"Integrador/src/sensor_turbuidez/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) domain.TurbiditySensor {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(sensor entities.TurbiditySensor) (entities.TurbiditySensor, error) {
	if sensor.Timestamp.IsZero() {
		sensor.Timestamp = time.Now()
	}

	query := `INSERT INTO turbidity_sensor (user_id, timestamp, turbidity) VALUES (?, ?, ?)`
	_, err := m.conn.Exec(query, sensor.UserID, sensor.Timestamp, sensor.Turbidity)
	if err != nil {
		return entities.TurbiditySensor{}, fmt.Errorf("error saving turbidity sensor data: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetByID(id, userID int) (entities.TurbiditySensor, error) {
	var sensor entities.TurbiditySensor
	query := `SELECT measurement_id, user_id, timestamp, turbidity FROM turbidity_sensor WHERE measurement_id = ? AND user_id = ?`
	err := m.conn.QueryRow(query, id, userID).Scan(
		&sensor.MeasurementID,
		&sensor.UserID,
		&sensor.Timestamp,
		&sensor.Turbidity,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sensor, errors.New("record not found")
		}
		return sensor, fmt.Errorf("error retrieving measurement: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetAll(userID int) ([]entities.TurbiditySensor, error) {
	query := `SELECT measurement_id, user_id, timestamp, turbidity FROM turbidity_sensor WHERE user_id = ?`
	rows, err := m.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving records: %v", err)
	}
	defer rows.Close()

	var sensors []entities.TurbiditySensor
	for rows.Next() {
		var sensor entities.TurbiditySensor
		if err := rows.Scan(
			&sensor.MeasurementID,
			&sensor.UserID,
			&sensor.Timestamp,
			&sensor.Turbidity,
		); err != nil {
			return nil, fmt.Errorf("error scanning record: %v", err)
		}
		sensors = append(sensors, sensor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over records: %v", err)
	}

	return sensors, nil
}

func (m *MySQL) Delete(id, userID int) error {
	query := `DELETE FROM turbidity_sensor WHERE measurement_id = ? AND user_id = ?`
	_, err := m.conn.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("error deleting record: %v", err)
	}
	return nil
}
