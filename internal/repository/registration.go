package repository

import (
	"MyCRUD/internal/models"
	"context"
	"fmt"
)

func (u *UserRepo) Registration(c *models.User) (id int, err error) {
	query := fmt.Sprintf("INSERT INTO %s (username, email, password) VALUES ($1, $2, $3) RETURNING id", "Users")
	row := u.Connection.QueryRow(context.Background(), query, c.Username, c.Email, c.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
