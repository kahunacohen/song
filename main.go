// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/controllers"
	// "github.com/kahunacohen/songs/controllers"
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
	router.Static("/static", "./static")
	router.Use(ResponseFormatMiddleware)
	router.GET("/api/v1/users/:user_id/songs", controllers.GetSongs(conn))
	router.GET("/users/:user_id/songs", controllers.GetSongs(conn))
	router.GET("/api/v1/users/:user_id/songs/:song_id", controllers.GetSong(conn))
	router.GET("/users/:user_id/songs/:song_id", controllers.GetSong(conn))
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
