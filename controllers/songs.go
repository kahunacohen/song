package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SongsByUserHandler(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		songID := c.Param("song_id")
		fmt.Println(db)
		if c.MustGet("rsp_fmt") == "json" {
			response := gin.H{
				"user_id": userID,
				"song_id": songID,
			}
			// Respond with the data
			c.JSON(http.StatusOK, response)
		} else {
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf("<ul><li>user: %s</li><li>song: %s</li></ul>", userID, songID)))
		}
	}
}
