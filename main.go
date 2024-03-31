// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songctls/controllers"
	"github.com/kahunacohen/songctls/models"
	"github.com/kahunacohen/songs/responders"
	"github.com/kahunacohen/songs/templates"
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
	router.GET("hello", func(c *gin.Context) {
		templates.Render(c, templates.Hello())
	})

	// /songs/
	router.GET("/users/:user_id/songs", controllers.ListSongs(conn, responders.SongList))
	router.PUT("/users/:user_id/songs/:song_id", controllers.UpdateSong(conn, responders.UpdateSong))
	router.GET("/users/:user_id/songs/new", func(c *gin.Context) {
		templates.NewSong(c.Param("user_id")).Render(c, c.Writer)
	})
	router.POST("/users/:user_id/songs/new", func(c *gin.Context) {
		var song *models.Song
		if err := c.ShouldBind(&song); err != nil {
			fmt.Println(err)
			return
		}
		if err := models.CreateSong(conn, song); err != nil {
			fmt.Println(err)
			return
		}
		c.Header("HX-Redirect", fmt.Sprintf("/users/%s/songs?flashOn=true&flashMsg=Song%%20added", c.Param("user_id")))
	})
	router.POST("/users/:user_id/songs/:song_id", controllers.UpdateSong(conn, responders.UpdateSong))
	router.GET("/users/:user_id/songs/:song_id", controllers.ReadSong(conn, responders.ReadSong))
	router.DELETE("/users/:user_id/songs/:song_id", func(c *gin.Context) {
		songID := c.Param("song_id")
		songAsInt, err := strconv.Atoi(songID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError,
				fmt.Errorf("not able to convert song id to an int: %v", err))
		}
		if models.DeleteSong(conn, songAsInt); err != nil {
			c.AbortWithError(http.StatusInternalServerError,
				fmt.Errorf("not able to delete song: %v", err))
		}
		c.Header("HX-Redirect", fmt.Sprintf("/users/%s/songs?flashOn=true&flashMsg=Song%%20deleted", c.Param("user_id")))
	})
	router.GET("/users/:user_id/artists", controllers.ListArtists(conn, responders.ArtistList))
	router.GET("/users/:user_id/artists/new", func(c *gin.Context) {
		templates.NewArtist(c.Param("user_id")).Render(c, c.Writer)
	})
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
