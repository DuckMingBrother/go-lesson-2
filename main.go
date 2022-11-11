package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var json Login

		if err := c.Bind(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "123" || json.Password != "123" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "帳號密碼錯誤"})
			return
		}

		c.JSON(http.StatusOK, "This is your token")
	})
	r.Run()
}

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
