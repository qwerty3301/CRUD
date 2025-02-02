package handlers

import (
	"MyCRUD/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		errorResponse(c, http.StatusBadRequest, "Empty Authorization header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		errorResponse(c, http.StatusBadRequest, "Invalid Input Data")
		return

	}
	userId, err := service.ParseToken(headerParts[1])
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid Token")
		return
	}
	c.Set("userId", userId)
}
