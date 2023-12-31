package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Song struct {
	CreatedAt time.Time `json:"created_at"`
	Capo      int
	Genre     string    `json:"genre"`
	SongID    int       `json:"song_id"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetSongsByUser(conn *pgx.Conn, userID int) ([]Song, error) {
	query := "SELECT song_id, title, genre FROM songs WHERE user_id = $1;"
	rows, err := conn.Query(context.Background(), query, userID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	var songs []Song
	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.SongID, &song.Title, &song.Genre); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		songs = append(songs, song)
	}
	defer rows.Close()
	return songs, nil
}
