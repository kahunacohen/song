package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SongsByUser(c *gin.Context) {
	userID := c.Param("user_id")
	songID := c.Param("song_id")
	c.MustGet("db")
	// getSong(db, userID, songID)
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
