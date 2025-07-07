package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"Integrador/src/motor/domain"
	"Integrador/src/motor/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) domain.MotorFermentadorRepository {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(sensor entities.MotorFermentador) (entities.MotorFermentador, error) {
	if sensor.Timestamp.IsZero() {
		sensor.Timestamp = time.Now()
	}
	query := `INSERT INTO motor_fermentador (user_id, timestamp, rpm) VALUES (?, ?, ?)`
	_, err := m.conn.Exec(query,
		sensor.UserID,
		sensor.Timestamp,
		sensor.RPM,
	)
	if err != nil {
		return entities.MotorFermentador{}, fmt.Errorf("error saving motor data: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetByID(id, userID int) (entities.MotorFermentador, error) {
	var sensor entities.MotorFermentador
	query := `SELECT measurement_id, user_id, timestamp, rpm  FROM motor_fermentador WHERE measurement_id = ? AND user_id = ?`
	err := m.conn.QueryRow(query, id, userID).Scan(
		&sensor.MeasurementID,
		&sensor.UserID,
		&sensor.Timestamp,
		&sensor.RPM,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sensor, errors.New("record not found")
		}
		return sensor, fmt.Errorf("error retrieving measurement: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetAll(userID int) ([]entities.MotorFermentador, error) {
	query := `SELECT measurement_id, user_id, timestamp, motor_fermentador FROM alcohol_sensor WHERE user_id = ?`
	rows, err := m.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving records: %v", err)
	}
	defer rows.Close()

	var sensors []entities.MotorFermentador
	for rows.Next() {
		var sensor entities.MotorFermentador
		if err := rows.Scan(
			&sensor.MeasurementID,
			&sensor.UserID,
			&sensor.Timestamp,
			&sensor.RPM,
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
	query := `DELETE FROM motor_fermentador WHERE measurement_id = ? AND user_id = ?`
	_, err := m.conn.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("error deleting record: %v", err)
	}
	return nil
}
