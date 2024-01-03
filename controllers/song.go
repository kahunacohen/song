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

func SongByUserHandler(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		songID := c.Param("song_id")
		songIDAsInt, _ := strconv.Atoi(songID)
		song, err := db.GetSongByID(conn, songIDAsInt)
		if err != nil {
			fmt.Println("error")
		}
		if c.MustGet("rsp_fmt") == "json" {
			c.JSON(http.StatusOK, song)
		} else {
			templates.Render(c, templates.Base(
				song.Title,
				templates.Song(*song),
			))
		}
	}
}
