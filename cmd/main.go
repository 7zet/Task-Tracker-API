package main

import (
	"github.com/7zet/Task-Tracker-API/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	internal.LoadUserFromFile()
	internal.LoadTaskFromFile()

	router := gin.Default()

	router.Use(internal.BasicAuth())
	// Auth
	router.POST("/register", internal.Register)
	router.POST("/login", internal.Login)
	router.POST("/tasks", internal.CreateTask)
	router.GET("/tasks", internal.GetTasks)

	router.Run(":8080")

}