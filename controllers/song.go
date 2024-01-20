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

type SongForm struct {
	Id     string `form:"id" binding:"required"`
	Title  string `form:"title" binding:"required"`
	Lyrics string `form:"lyrics" binding:"required"`
}

func PutSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		var songForm SongForm
		c.Bind(&songForm)
		id, err := strconv.Atoi(songForm.Id)
		if err != nil {
			//
		}
		_, err = conn.Exec(c, "UPDATE songs SET title=$1, lyrics=$2 WHERE id=$3",
			songForm.Title, songForm.Lyrics, id)
		if err != nil {
			fmt.Println("error!")
		}

	}
}
