package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Song struct {
	CreatedAt time.Time `json:"created_at"`
	Capo      int
	Artist    string    `form:"artist" binding:"required" json:"artist"`
	Genre     string    `form:"genre" json:"genre"`
	Lyrics    string    `form:"lyrics" binding:"required" json:"lyrics"`
	Id        int       `form:"id" json:"id"`
	Title     string    `form:"title" binding:"required" json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int       `form:"user_id" json:"user_id"`
}

func CreateSong(conn *pgx.Conn, song *Song) error {
	fmt.Println("CREATING!!!")
	var id int
	query := "INSERT INTO songs (title, lyrics, user_id, genre_id, artist_id) VALUES($1, $2, $3, $4, $5) RETURNING id"
	err := conn.QueryRow(context.Background(), query, song.Title, song.Lyrics, song.UserID, 1, song.Artist).Scan(&id)
	if err != nil {
		return fmt.Errorf("error creating song: %v", err)
	}
	song.Id = id
	return nil
}
func DeleteSong(conn *pgx.Conn, songID int) error {
	query := "DELETE FROM songs WHERE id=$1"
	_, err := conn.Exec(context.Background(), query, songID)
	if err != nil {
		return err
	}
	return nil
}
