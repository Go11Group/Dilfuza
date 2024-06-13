package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Go11Group/Dilfuza/EXAM2/models"
	"net/http"
)

func (h *Handler) Course(c *gin.Context) {
	id := c.Param("id")
	course, err := h.Courses.CourseGet(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *Handler) CourseCreate(c *gin.Context) {
	var course models.Courses
	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Courses.CourseCreate(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)
}

func (h *Handler) CourseUpdate(c *gin.Context) {
	id := c.Param("id")
	var course models.Courses
	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course.CourseId = id
	if err := h.Courses.CourseUpdate(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *Handler) CourseDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Courses.CourseDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
