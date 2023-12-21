// main.go
package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kahunacohen/songs/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/api/v1/users/:user_id/:song_id", controllers.SongsByUser)
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
