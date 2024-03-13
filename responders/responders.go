package responders

import (
	"github.com/gin-gonic/gin"
	"github.com/kahunacohen/songctls/mdls"
	"github.com/kahunacohen/songs/templates"
)

func SongList(context *gin.Context, userID string, songs []mdls.Song, totalCount, page int, searchTerm string, partial bool) {
	if partial {
		templates.Render(context, templates.SongTable(songs, totalCount, page))
	} else {
		templates.Render(context, templates.Base(
			"My Songs",
			templates.Songs(userID, songs, totalCount, page, searchTerm),
		))
	}
}
