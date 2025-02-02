package handlers

import (
	"MyCRUD/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) Authorization(c *gin.Context) {
	var inputValue *models.AuthorizationData
	if err := c.BindJSON(&inputValue); err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid Input Data")
		return
	}

	token, err := h.Service.AuthorizationService(inputValue)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Authorization Error")
		return
	}
	c.JSON(200, map[string]interface{}{
		"token": token,
	})
}
