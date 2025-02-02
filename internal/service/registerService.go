package service

import (
	"MyCRUD/internal/models"
	"crypto/sha1"
	"fmt"
)

const salt = "adsfgs1fhdjhj1jdhsgas3dhm"

func (u *UserService) RegisterService(c *models.User) (int, error) {
	c.Password = GeneratePasswordHash(c.Password)
	return u.Repo.Registration(c)
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
