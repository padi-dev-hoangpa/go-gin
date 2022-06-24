package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/square/:number", func(c *gin.Context) {
		number := c.Param("number")
		i, err := strconv.ParseInt(number, 0, 32)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"status":  false,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"result": i * i,
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
