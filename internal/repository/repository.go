package repository

import (
	"MyCRUD/internal/models"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	GetUser(*models.AuthorizationData) (*models.User, error)
	Registration(*models.User) (int, error)
	Delete(id int) error
	Update(id int, newUsername string) error
}

type UserRepo struct {
	Connection *pgx.Conn
}
