package repository

import (
	"context"
	"fmt"
)

func (u *UserRepo) Update(userId int, newUsername string) error {
	query := fmt.Sprintf("UPDATE Users SET username = $1 WHERE id = $2;")
	_, err := u.Connection.Exec(context.Background(), query, newUsername, userId)
	if err != nil {
		return err
	}
	return nil
}
