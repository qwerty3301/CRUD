package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewUsernameData struct {
	NewUsername string `json:"newUsername" binding:"required"`
}

func (h *UserHandler) UpdateHandler(c *gin.Context) {
	var inputData NewUsernameData

	if err := c.BindJSON(&inputData); err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid Input Data")
	}
	id, ok := c.Get("userId")
	if !ok {
		errorResponse(c, http.StatusBadRequest, "Update Failed")
		return
	}

	err := h.Service.UpdateService(id.(int), inputData.NewUsername)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Update Failed")
		return
	}
}
