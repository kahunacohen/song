package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/db"
	"github.com/kahunacohen/songs/templates"
)

func GetSongs(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		userIDAsInt, _ := strconv.Atoi(userID)
		songs, err := db.GetSongsByUser(conn, userIDAsInt)
		if err != nil {
			fmt.Println("error")
		}
		if c.MustGet("rsp_fmt") == "json" {
			c.JSON(http.StatusOK, songs)
		} else {
			templates.Render(c, templates.Base(
				"My Songs",
				templates.Songs(userID, songs),
			))
		}
	}
}

func NewSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		templates.NewSong(c.Param("user_id")).Render(c, c.Writer)
	}
}
func PostSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		var song db.Song
		// Bind form data to the User struct
		if err := c.ShouldBind(&song); err != nil {
			fmt.Println(err)
			return
		}
		err := db.CreateSong(conn, &song)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("genre:")
		fmt.Println(song.Genre)
		templates.SongRow(song).Render(c, c.Writer)
	}
}
func DeleteSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		songID := c.Param("song_id")
		songAsInt, err := strconv.Atoi(songID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError,
				fmt.Errorf("not able to convert song id to an int: %v", err))
		}
		if db.DeleteSong(conn, songAsInt); err != nil {
			c.AbortWithError(http.StatusInternalServerError,
				fmt.Errorf("not able to delete song: %v", err))
		}
		c.Header("HX-Redirect", fmt.Sprintf("/users/%s/songs?confirmation=deleted", c.Param("user_id")))
	}
}
