package main

import (
	"log"
	"net/http"

	model "github.com/Go11Group/Dilfuza/Gin/models"
	"github.com/Go11Group/Dilfuza/Gin/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()

	// User routes
	r.POST("/user/create", CreateUser)
	r.PUT("/user/update", UpdateUser)
	r.DELETE("/user/delete", DeleteUser)

	// Problem routes
	r.POST("/problem/create", CreateProblem)
	r.PUT("/problem/update", UpdateProblem)
	r.DELETE("/problem/delete", DeleteProblem)

	log.Println("Server is running...")
	if err := r.Run("localhost:8088"); err != nil {
		log.Println("Error server is running!")
		return
	}
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		log.Println("error while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id := uuid.NewString()
	user.Id = id

	respUser, err := storage.CreateUser(&user)
	if err != nil {
		log.Println("error while creating user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, respUser)
}

func UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		log.Println("error while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userID := c.Query("id")

	respUser, err := storage.UpdateUser(userID, &user)
	if err != nil {
		log.Println("error while updating user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, respUser)
}

func DeleteUser(c *gin.Context) {
	userID := c.Query("id")

	if err := storage.DeleteUser(userID); err != nil {
		log.Println("error while deleting user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted User"})
}

func CreateProblem(c *gin.Context) {
	var problem model.Problem
	if err := c.BindJSON(&problem); err != nil {
		log.Println("Error while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id := uuid.NewString()
	problem.Id = id

	respProblem, err := storage.CreateProblem(&problem)
	if err != nil {
		log.Println("Error while creating problem", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create problem"})
		return
	}

	c.JSON(http.StatusCreated, respProblem)
}

func UpdateProblem(c *gin.Context) {
	var problem model.Problem
	if err := c.BindJSON(&problem); err != nil {
		log.Println("Error while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	problemId := c.Query("id")

	respProblem, err := storage.UpdateProblem(problemId, &problem)
	if err != nil {
		log.Println("Error while updating problem", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update problem"})
		return
	}

	c.JSON(http.StatusOK, respProblem)
}

func DeleteProblem(c *gin.Context) {
	problemId := c.Query("id")

	if err := storage.DeleteProblem(problemId); err != nil {
		log.Println("Error while deleting problem", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete problem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Problem"})
}
