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
		userID := c.Param("user_id")
		songID := c.Param("song_id")
		userIDAsInt, _ := strconv.Atoi(userID)
		songIDAsInt, _ := strconv.Atoi(songID)
		songs, err := db.GetSongByID(conn, songIDAsInt)
		if err != nil {
			fmt.Println("error")
		}
		if c.MustGet("rsp_fmt") == "json" {
			c.JSON(http.StatusOK, songs)
		} else {
			templates.Render(c, templates.Base(
				"Songs",
				templates.Songs(userID, songs),
			))
		}
	}
}
