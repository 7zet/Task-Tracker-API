package internal

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"

)

func BasicAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		data, _ := os.ReadFile("users.json")
		var users []User

		json.Unmarshal(data, &users)
		
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatus(401)
			return 
		}

		for _, u := range users {
			if u.Username == username && u.Password == password{
				c.Next()
				return 
			}
		}
		c.AbortWithStatus(401)
	}
}