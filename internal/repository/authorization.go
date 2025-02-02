package repository

import (
	"MyCRUD/internal/models"
	"context"
	"fmt"
)

func (u *UserRepo) GetUser(data *models.AuthorizationData) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf("select id, username, email, password from Users  where email = $1 and password = $2")
	row := u.Connection.QueryRow(context.Background(), query, data.Email, data.Password)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
