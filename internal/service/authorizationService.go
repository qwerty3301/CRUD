package service

import (
	"MyCRUD/internal/models"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

const JWTsecret = "hello"

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"id"`
}

func (u *UserService) AuthorizationService(data *models.AuthorizationData) (string, error) {
	data.Password = GeneratePasswordHash(data.Password)
	user, err := u.Repo.GetUser(data)
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(user)
	return token, err
}

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, UserId: user.ID,
	})
	return token.SignedString([]byte(JWTsecret))
}

func ParseToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(JWTsecret), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("bad token")
	}
	return claims.UserId, nil
}
