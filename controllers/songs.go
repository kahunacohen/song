package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/mdls"
	"github.com/kahunacohen/songs/templates"
)

func ListSongs(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		userIDAsInt, _ := strconv.Atoi(userID)
		page := c.Query("page")
		content := c.Query("ct")
		q := c.Query("q")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			pageInt = 1
		}
		queries := mdls.New(conn)
		var songs []mdls.SongByUser
		params := mdls.GetSongsByUserParams{UserID: int32(userIDAsInt), Offset: int32(pageInt - 1)}
		if q != "" {
			// Column2 is the search query string due to some weird sqlc issue not properly naming
			// the column.
			params.Column2 = q
		}
		songs, err = queries.GetSongsByUser(c, params)
		if err != nil {
			fmt.Println(err)
			return
		}
		totalCount, err := queries.GetTotalSongsByUser(c, int32(userIDAsInt))
		if err != nil {
			fmt.Println("error getting total count")
		}
		if content == "partial" {
			templates.Render(c, templates.SongTable(songs, int(totalCount), pageInt))
		} else {
			templates.Render(c, templates.Base(
				"My Songs",
				templates.Songs(userID, songs, int(totalCount), pageInt, c.Query("q")),
			))
		}
	}
}

func SearchSongs(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		userIDAsInt, _ := strconv.Atoi(userID)
		page := c.Query("page")
		content := c.Query("ct")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			pageInt = 1
		}
		queries := mdls.New(conn)
		songs, err := queries.GetSongsByUser(c, mdls.GetSongsByUserParams{UserID: int32(userIDAsInt), Offset: int32(pageInt - 1)})
		if err != nil {
			fmt.Println("error getting songs")
		}
		totalCount, err := queries.GetTotalSongsByUser(c, int32(userIDAsInt))
		if err != nil {
			fmt.Println("error getting total count")
		}
		if content == "partial" {
			templates.Render(c, templates.SongTable(songs, int(totalCount), pageInt))
		} else {
			templates.Render(c, templates.Base(
				"My Songs",
				templates.Songs(userID, songs, int(totalCount), pageInt, c.Query("q")),
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
		song, _ := queries.GetSongByUser(c, int32(songIDAsInt))
		uri := fmt.Sprintf("/users/%s/songs/%d", userID, song.SongID)
		editModeUri := fmt.Sprintf("%s?mode=edit", uri)
		mode := c.Query("mode")
		artists, err := queries.GetArtistsByUser(
			c, mdls.GetArtistsByUserParams{UserID: int32(userIdAsInt), Offset: 0})
		if err != nil {
			fmt.Println(err)
		}
		templates.Render(c, templates.Base(
			func() string {
				if mode == "edit" {
					return fmt.Sprintf("Edit \"%s\"", song.SongTitle)
				} else {
					return song.SongTitle
				}
			}(),
			templates.Song(song, artists, uri, editModeUri, mode == "edit"),
		))
	}
}

func UpdateSong(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		var song mdls.SongByUser
		c.Bind(&song)
		fmt.Printf("title: %s, lyrics: %s, id: %d\n", song.SongTitle, song.SongLyrics, song.SongID)
		song.SongTitle = SanitizeInput(song.SongTitle)
		song.SongLyrics = SanitizeInput(song.SongLyrics)
		fmt.Println(song.SongID)
		// err := models.UpdateSong(conn, song)
		queries := mdls.New(conn)
		err := queries.UpdateSong(c, mdls.UpdateSongParams{Title: song.SongTitle, Lyrics: song.SongLyrics, ID: song.SongID})
		if err != nil {
			// @TODO error handling.
			fmt.Printf("error updating song: %v\n", err)
		}
		fmt.Println(song.SongID)
		uri := fmt.Sprintf("/users/%s/songs/%d?flashOn=true&flashMsg=Song%%20saved", userID, song.SongID)
		if c.Request.Method == "POST" {
			// We are receiving from old-school form where method=POST
			// is not supported by browsers, so redirect to same page
			// with a GET.
			c.Redirect(http.StatusSeeOther, uri)
			return
		}
		// var artists []models.Artist
		// templates.Render(c, templates.SongFormContents(song, artists))
	}
}
