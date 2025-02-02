package handlers

import (
	"MyCRUD/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) RegisterHandler(c *gin.Context) {
	var inputValue models.User

	if err := c.BindJSON(&inputValue); err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid Input Data")
		return
	}

	answer, err := h.Service.RegisterService(&inputValue)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Registration Error")
		return
	}
	c.JSON(200, answer)
}
