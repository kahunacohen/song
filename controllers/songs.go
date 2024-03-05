package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/models"
	"github.com/kahunacohen/songs/templates"
)

func ReadSongs(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		userIDAsInt, _ := strconv.Atoi(userID)
		offsetParam := c.Query("offset")
		offset, err := strconv.Atoi(offsetParam)
		if err != nil {
			offset = 0
		}
		songs, err := models.GetSongsByUser(conn, userIDAsInt, offset)
		if err != nil {
			fmt.Println(err)
		}
		if c.MustGet("rsp_fmt") == "json" {
			c.JSON(http.StatusOK, songs)
		} else {
			templates.Render(c, templates.Base(
				"My Songs",
				templates.Songs(userID, songs, c.Query("q")),
			))
		}
	}
}

func SearchSongs(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		userIDAsInt, _ := strconv.Atoi(userID)
		q := c.Query("q")
		songs, err := models.SearchSongs(conn, userIDAsInt, q)
		if err != nil {
			fmt.Println(err)
		}
		templates.Render(c, templates.SongRows(songs))
	}
}

func NewSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		templates.NewSong(c.Param("user_id")).Render(c, c.Writer)
	}
}
func CreateSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		var song *models.Song
		if err := c.ShouldBind(&song); err != nil {
			fmt.Println(err)
			return
		}
		if err := models.CreateSong(conn, song); err != nil {
			fmt.Println(err)
			return
		}
		templates.SongRow(*song).Render(c, c.Writer)
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
		if models.DeleteSong(conn, songAsInt); err != nil {
			c.AbortWithError(http.StatusInternalServerError,
				fmt.Errorf("not able to delete song: %v", err))
		}
		c.Header("HX-Redirect", fmt.Sprintf("/users/%s/songs?flashOn=true&flashMsg=Song%%20deleted", c.Param("user_id")))
	}
}
