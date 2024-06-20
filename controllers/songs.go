package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/mdls"
	"github.com/kahunacohen/songs/models"
	"github.com/kahunacohen/songs/templates"
	"github.com/kahunacohen/songs/wrapped_models"
)

func ListSongs(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("LIST SONGS")
		userID := c.Param("user_id")
		userIDAsInt, _ := strconv.Atoi(userID)
		q := c.Query("q")
		page := c.Query("page")
		content := c.Query("ct")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			pageInt = 1
		}
		songs, totalCount, err := models.SearchSongs(conn, userIDAsInt, q, pageInt)
		if err != nil {
			fmt.Println(err)
		}
		if content == "partial" {
			fmt.Println("PARTIAL")
			templates.Render(c, templates.SongTable(songs, totalCount, pageInt))
		} else {
			templates.Render(c, templates.Base(
				"My Songs",
				templates.Songs(userID, songs, totalCount, pageInt, c.Query("q")),
			))
		}
	}
}

func ReadSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		songID := c.Param("song_id")
		songIDAsInt, _ := strconv.Atoi(songID)
		userID := c.Param("user_id")
		userIdAsInt, _ := strconv.Atoi(userID)
		queries := mdls.New(conn)
		song, _ := queries.GetSong(c, int32(songIDAsInt))
		wrappedSong := wrapped_models.New(song)
		uri := fmt.Sprintf("/users/%s/songs/%d", userID, song.SongID)
		editModeUri := fmt.Sprintf("%s?mode=edit", uri)
		mode := c.Query("mode")
		artists, _, _ := models.SearchArtists(conn, userIdAsInt, nil, nil)
		templates.Render(c, templates.Base(
			func() string {
				if mode == "edit" {
					return fmt.Sprintf("Edit \"%s\"", song.Title)
				} else {
					return song.Title
				}
			}(),
			templates.Song(wrappedSong, artists, uri, editModeUri, mode == "edit"),
		))
	}
}

func UpdateSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("UPDATE")
		userID := c.Param("user_id")
		var song mdls.Song
		c.Bind(&song)
		song.Title = SanitizeInput(song.Title)
		song.Lyrics = SanitizeInput(song.Lyrics)
		// err := models.UpdateSong(conn, song)
		queries := mdls.New(conn)
		err := queries.UpdateSong(c, mdls.UpdateSongParams{Title: song.Title, Lyrics: song.Lyrics, ID: song.ID})
		if err != nil {
			// @TODO error handling.
			fmt.Printf("error updating song: %v\n", err)

		}
		uri := fmt.Sprintf("/users/%s/songs/%d?flashOn=true&flashMsg=Song%%20saved", userID, song.ID)
		if c.Request.Method == "POST" {
			// We are receiving from old-school form where method=POST
			// is not supported by browsers, so redirect to same page
			// with a GET.
			c.Redirect(http.StatusSeeOther, uri)
			return
		}
		var artists []models.Artist
		templates.Render(c, templates.SongFormContents(song, artists))
	}
}
