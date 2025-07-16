package adapters

import (
	"Integrador/src/users/domain/entities"
	"database/sql"
	"errors"
	"fmt"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(user entities.User) error {
	query := `INSERT INTO user (name, lastName, email, password, id_rol, code) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := m.conn.Exec(query, user.Name, user.LastName, user.Email, user.Password, user.Role, user.Code)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}
	return nil
}

func (m *MySQL) GetByEmail(email string) (entities.User, error) {
	var user entities.User
	query := `SELECT id_user, name, lastName, email, password, id_rol, code FROM user WHERE email = ? LIMIT 1`

	err := m.conn.QueryRow(query, email).Scan(
		&user.ID, &user.Name, &user.LastName, &user.Email, &user.Password, &user.Role, &user.Code,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, fmt.Errorf("failed to retrieve user by email: %v", err)
	}
	return user, nil
}

func (m *MySQL) GetAll() ([]entities.User, error) {
	query := "SELECT id_user, name, lastName, email, password, id_rol, code FROM user"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %v", err)
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Password, &user.Role, &user.Code)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return users, nil
}

func (m *MySQL) GetById(id int) (entities.User, error) {
	query := "SELECT id_user, name, lastName, email, password, id_rol, code FROM user WHERE id_user = ?"
	row := m.conn.QueryRow(query, id)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Password, &user.Role, &user.Code)
	if err == sql.ErrNoRows {
		return entities.User{}, errors.New("user not found")
	} else if err != nil {
		return entities.User{}, fmt.Errorf("failed to retrieve user: %v", err)
	}

	return user, nil
}

func (m *MySQL) Edit(user entities.User) error {
	query := "UPDATE user SET name = ?, lastName = ?, email = ?, password = ?, id_rol = ?, code = ? WHERE id_user = ?"
	_, err := m.conn.Exec(query, user.Name, user.LastName, user.Email, user.Password, user.Role, user.Code, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM user WHERE id_user = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
