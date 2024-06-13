package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Go11Group/Dilfuza/EXAM2/models"
	"net/http"
)

func (h *Handler) User(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Users.UserGet(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) UserCreate(c *gin.Context) {
	var user models.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	if err := h.Users.UserCreate(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) UserUpdate(c *gin.Context) {
	id := c.Param("id")
	var user models.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.UserId = id
	if err := h.Users.UserUpdate(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (h *Handler) UserDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Users.UserDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
