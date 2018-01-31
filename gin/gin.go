package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/user/:id", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"User": c.Param("id"),
			"Name": "Yamada",
		})
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8888
}
