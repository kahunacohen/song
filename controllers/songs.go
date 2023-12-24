package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/db"
)

func SongsByUserHandler(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		userIDAsInt, _ := strconv.Atoi(userID)
		songs, err := db.GetSongsByUser(conn, userIDAsInt)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(songs)

		if c.MustGet("rsp_fmt") == "json" {
			response := gin.H{
				"user_id": userID,
				"song_id": "foo",
			}
			c.JSON(http.StatusOK, response)
		} else {
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf("<ul><li>user: %s</li><li>song: %s</li></ul>", userID, "foo")))
		}
	}
}
