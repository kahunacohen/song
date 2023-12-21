// main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kahunacohen/songs/controllers"
)

func ResponseFormatMiddleware(c *gin.Context) {
	if strings.Contains(c.Request.URL.Path, "/api") {
		c.Set("rsp_fmt", "json")
	} else {
		c.Set("rsp_fmt", "html")
	}
	c.Next()
}
func main() {
	router := gin.Default()
	router.Use(ResponseFormatMiddleware)
	router.GET("/api/v1/users/:user_id/:song_id", controllers.SongsByUser)
	router.GET("/users/:user_id/:song_id", controllers.SongsByUser)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
