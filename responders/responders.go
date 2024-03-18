package responders

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kahunacohen/songctls/models"
	"github.com/kahunacohen/songs/templates"
)

func SongList(context *gin.Context, userID string, songs []models.Song, totalCount, page int, searchTerm string, partial bool) {
	if partial {
		templates.Render(context, templates.SongTable(songs, totalCount, page))
	} else {
		templates.Render(context, templates.Base(
			"My Songs",
			templates.Songs(userID, songs, totalCount, page, searchTerm),
		))
	}
}

func ReadSong(context *gin.Context, mode string, song models.Song, uri string, editModeUri string) {
	fmt.Println("RESPOND")
	templates.Render(context, templates.Base(
		func() string {
			if mode == "edit" {
				return fmt.Sprintf("Edit \"%s\"", song.Title)
			} else {
				return song.Title
			}
		}(),
		templates.Song(song, uri, editModeUri, mode == "edit"),
	))
}

func UpdateSong(c *gin.Context, song models.Song) {
	templates.Render(c, templates.SongFormContents(song))
}
