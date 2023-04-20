package db

import "github.com/google/uuid"

func (d *Database) GetUserID(username, password string) (*uuid.UUID, error) {
	query := "SELECT id FROM users WHERE username = $1 and password = $2"
	var id uuid.UUID
	err := d.Client.QueryRow(query, username, password).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
}
