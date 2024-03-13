package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	//foo "github.com/kahunacohen/song-controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/models"
	"github.com/kahunacohen/songs/templates"
)

func ListSongs(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		userIDAsInt, _ := strconv.Atoi(userID)
		q := c.Query("q")
		page := c.Query("page")
		content := c.Query("ct")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			pageInt = 1
		}
		songs, totalCount, err := models.SearchSongs(conn, userIDAsInt, q, pageInt)
		if err != nil {
			fmt.Println(err)
		}

		if content == "partial" {
			fmt.Printf("page: %d\n", pageInt)

			templates.Render(c, templates.SongTable(songs, totalCount, pageInt))
		} else {
			templates.Render(c, templates.Base(
				"My Songs",
				templates.Songs(userID, songs, totalCount, pageInt, c.Query("q")),
			))
		}
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
		c.Header("HX-Redirect", fmt.Sprintf("/users/%s/songs?flashOn=true&flashMsg=Song%%20added", c.Param("user_id")))

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
