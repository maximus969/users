package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maximus969/users-app"
)

func (h *Handler) createUser(c *gin.Context) {
	var input users.User	

	if err:= c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Users.Create(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User was created",
	})
}

func (h *Handler) getAllUsers(c *gin.Context) {
	//
}

func (h *Handler) getUser(c *gin.Context) {
	//
}

func (h *Handler) updateUser(c *gin.Context) {
	//
}

func (h *Handler) deleteUser(c *gin.Context) {
	//
}