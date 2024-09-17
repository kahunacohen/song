package utils

import (
	"fmt"
	"regexp"

	"github.com/kahunacohen/songs/mdls"
)

func SongToHTML(song mdls.SongByUser) string {
	var ret string
	chordRe := regexp.MustCompile(`\[(.+?)\]`)
	ret = fmt.Sprintf(
		"<div style='position: relative;'>%s",
		chordRe.ReplaceAllString(song.SongLyrics, "<span style='position:absolute;top:-12px;font-size:90%;font-weight:bold;'>$1</span>"))
	lineWrapperRe := regexp.MustCompile(`(?m)^.*$`)
	ret = lineWrapperRe.ReplaceAllString(ret, "<div style='position:relative;line-height:2.5;'>$0</div>")
	return ret
}
