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
	Title  string `form:"title" binding:"required"`
	Lyrics string `form:"lyrics" binding:"required"`
}

func PutSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var request PutSongRequest
		// if err := c.BindJSON(&request); err != nil {
		// 	// DO SOMETHING WITH THE ERROR
		// }

		var songForm SongForm

		// Bind and validate form data
		if err := c.ShouldBind(&songForm); err != nil {
			fmt.Println("error!")
			return
		}
		fmt.Println(songForm.Title)

	}
}
