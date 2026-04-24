package main

import (
	"github.com/7zet/Task-Tracker-API/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	internal.LoadUserFromFile()

	router := gin.Default()

	// Auth
	router.POST("/register", internal.Register)
	router.POST("/login", internal.Login)
	router.POST("/tasks", internal.CreateTask)
	router.GET("/tasks", internal.GetTasks)

	router.Run(":8080")

}