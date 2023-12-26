// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/controllers"
)

func initDB(ctx context.Context) (*pgx.Conn, error) {
	connStr := "postgresql://postgres:password@postgres:5432/songs?sslmode=disable"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn, err
}
func main() {
	ctx := context.Background()
	conn, err := initDB(ctx)
	if err != nil {
		log.Fatalf("failed to intialize db: %v", err)
	}
	defer conn.Close(ctx)
	router := gin.Default()
	router.Use(ResponseFormatMiddleware)
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/api/v1/users/:user_id/:song_id", controllers.SongsByUserHandler(conn))
	router.GET("/users/:user_id/:song_id", controllers.SongsByUserHandler(conn))
	router.GET("/", func(c *gin.Context) {
		data := gin.H{
			"Title": "Gin Example",
		}

		// Render the HTML template with the provided data
		c.HTML(http.StatusOK, "base.html", data)
	})

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
