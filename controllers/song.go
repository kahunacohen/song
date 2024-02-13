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

func ReadSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		songID := c.Param("song_id")
		songIDAsInt, _ := strconv.Atoi(songID)
		userID := c.Param("user_id")
		song, getSongErr := models.GetSongByID(conn, songIDAsInt)
		if c.MustGet("rsp_fmt") == "json" {
			c.JSON(http.StatusOK, song)
		} else {
			if getSongErr != nil {
				templates.Render(c, templates.Base("Not found", templates.NotFound()))
				return
			}
			uri := fmt.Sprintf("/users/%s/songs/%d", userID, song.Id)
			editModeUri := fmt.Sprintf("%s?mode=edit", uri)
			mode := c.Query("mode")
			templates.Render(c, templates.Base(
				func() string {
					if mode == "edit" {
						return fmt.Sprintf("Edit \"%s\"", song.Title)
					} else {
						return song.Title
					}
				}(),
				templates.Song(*song, uri, editModeUri, mode == "edit"),
			))
		}
	}
}

func UpdateSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		var song models.Song
		c.Bind(&song)
		_, err := conn.Exec(c, "UPDATE songs SET title=$1, lyrics=$2 WHERE id=$3",
			song.Title, song.Lyrics, song.Id)
		if err != nil {
			// @TODO error handling.
			fmt.Println("error!")
		}
		if c.MustGet("rsp_fmt") == "json" {
			c.JSON(http.StatusOK, song)
		} else {
			uri := fmt.Sprintf("/users/%s/songs/%d?confirmation=saved", userID, song.Id)
			if c.Request.Method == "POST" {
				// We are receiving from old-school form where method=POST
				// is not supported by browsers, so redirect to same page
				// with a GET.
				c.Redirect(http.StatusSeeOther, uri)
				return
			}
			// Re-render the new data when PUTing
			templates.Render(c, templates.SongFormContents(song))
		}
	}
}
