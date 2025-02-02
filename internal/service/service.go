package service

import (
	"MyCRUD/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}
