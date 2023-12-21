package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SongsByUser(c *gin.Context) {
	userID := c.Param("user_id")
	songID := c.Param("song_id")

	response := gin.H{
		"user_id": userID,
		"song_id": songID,
	}
	// Respond with the data
	c.JSON(http.StatusOK, response)
}
