package repository

import (
	"context"
	"errors"
	"fmt"
)

func (u *UserRepo) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM Users WHERE Id=$1")
	_, err := u.Connection.Exec(context.Background(), query, userId)
	if err != nil {
		return errors.New("Error deleting user")
	}
	return nil
}
