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
	Genre     string    `form:"genre" json:"genre"`
	Lyrics    string    `form:"lyrics" binding:"required" json:"lyrics"`
	Id        int       `form:"id" json:"id"`
	Title     string    `form:"title" binding:"required" json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int       `form:"user_id" json:"user_id"`
}

func GetSongsByUser(conn *pgx.Conn, userID int) ([]Song, error) {
	query := "SELECT id, title, genre, user_id FROM songs WHERE user_id = $1 ORDER BY title;"
	rows, err := conn.Query(context.Background(), query, userID)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	var songs []Song
	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.Id, &song.Title, &song.Genre, &song.UserID); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		songs = append(songs, song)
	}
	defer rows.Close()
	return songs, nil
}

func GetSongByID(conn *pgx.Conn, id int) (*Song, error) {
	query := "SELECT id, title, genre, lyrics FROM songs WHERE id = $1;"
	row := conn.QueryRow(context.Background(), query, id)
	var song Song
	if err := row.Scan(&song.Id, &song.Title, &song.Genre, &song.Lyrics); err != nil {
		return nil, fmt.Errorf("error scanning row: %v", err)
	}
	return &song, nil
}
func UpdateSong(conn *pgx.Conn, song *Song) error {
	query := "UPDATE songs SET title='$1', lyrics='$2' WHERE id=$3;"
	_, err := conn.Exec(context.Background(), query, song.Title, song.Lyrics, song.Id)
	if err != nil {
		return fmt.Errorf("error updating song: %v", err)
	}
	return nil
}
func CreateSong(conn *pgx.Conn, song *Song) error {
	var id int
	query := "INSERT INTO songs (title, lyrics, user_id, genre) VALUES($1, $2, $3, $4) RETURNING id"
	err := conn.QueryRow(context.Background(), query, song.Title, song.Lyrics, song.UserID, "Rock").Scan(&id)
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
		fmt.Println("error here!")
		return err
	}
	return nil
}
