// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kahunacohen/songs/controllers"
)

func initDB() (*pgxpool.Pool, error) {
	connStr := "postgresql://postgres:password@postgres:5432/songs?sslmode=disable"
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("failed to intialize db: %v", err)
	}
	defer db.Close()
	router := gin.Default()
	router.Use(ResponseFormatMiddleware)
	router.Use(DatabaseMiddleware(db))
	router.GET("/api/v1/users/:user_id/:song_id", controllers.SongsByUser)
	router.GET("/users/:user_id/:song_id", controllers.SongsByUser)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
