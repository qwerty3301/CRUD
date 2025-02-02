package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) DeleteHandler(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errorResponse(c, http.StatusBadRequest, "Delete Error")
		return
	}
	err := h.Service.DeleteService(id.(int))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Delete Error: "+err.Error())
		return
	}
}
