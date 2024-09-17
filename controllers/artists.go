package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/mdls"
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
		artistId := c.Param("artist_id")
		artistIdAsInt, _ := strconv.Atoi(artistId)
		queries := mdls.New(conn)
		if c.Request.Method == "GET" {
			artistName, err := queries.GetArtist(c, int32(artistIdAsInt))
			if err != nil {
				fmt.Println("error getting artist")
				return
			}

			templates.Render(c, templates.Base(artistName, templates.Artist(c.Request.RequestURI, artistName)))
		} else { // POST
			artistName := c.PostForm("artist")
			err := queries.UpdateArtist(c, mdls.UpdateArtistParams{Name: artistName, ID: int32(artistIdAsInt)})
			if err != nil {
				fmt.Println("error updating artist")
			}
			c.Header("HX-Redirect", fmt.Sprintf("/users/%s/artists?flashOn=true&flashMsg=Artist%%20updated", c.Param("user_id")))
		}

	}
}
