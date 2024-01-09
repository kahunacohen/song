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

func GetSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		songID := c.Param("song_id")
		songIDAsInt, _ := strconv.Atoi(songID)
		userID := c.Param("user_id")
		song, err := db.GetSongByID(conn, songIDAsInt)
		if err != nil {
			fmt.Println("error")
		}
		if c.MustGet("rsp_fmt") == "json" {
			c.JSON(http.StatusOK, song)
		} else {
			uri := fmt.Sprintf("/users/%s/songs/%d", userID, song.Id)
			uriEditMode := fmt.Sprintf("%s?mode=edit", uri)
			templates.Render(c, templates.Base(
				song.Title,
				templates.Song(*song, uri, uriEditMode, c.Query("mode") == "edit"),
			))
		}
	}
}
