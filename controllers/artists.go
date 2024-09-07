package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/models"
	"github.com/kahunacohen/songs/templates"
)

func ListArtists(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		userIDAsInt, _ := strconv.Atoi(userID)
		q := c.Query("q")
		page := c.Query("page")
		content := c.Query("ct")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			pageInt = 1
		}
		artists, totalCount, err := models.SearchArtists(conn, userIDAsInt, &q, &pageInt)
		if err != nil {
			fmt.Println(err)
		}
		partial := content == "partial"
		if partial {
			templates.Render(c, templates.ArtistTable(artists, totalCount, pageInt))
		} else {
			templates.Render(c, templates.Base(
				"My Artists",
				templates.Artists(userID, artists, totalCount, pageInt, c.Query("q")),
			))
		}
	}
}

func UpdateArtist(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
