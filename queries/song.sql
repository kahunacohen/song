-- name: GetSongByUser :one
SELECT * FROM songs_by_user WHERE song_id = $1;

-- name: GetSongsByUser :many
SELECT 
    user_id, 
    user_first_name, 
    user_last_name, 
    user_email, 
    song_id, 
    song_title, 
    song_lyrics, 
    genre_id, 
    genre, 
    artist_id, 
    artist_name 
FROM 
    songs_by_user 
WHERE 
    user_id = $1 
    AND (
        $2::text IS NULL 
        OR $2 = '' 
        OR CONCAT(song_title, ' ', artist_name) ILIKE '%' || $2::text || '%'
    )
ORDER BY 
    song_title 
LIMIT 
    10 
OFFSET 
    $3;

-- name: GetTotalSongsByUser :one
SELECT COUNT(*) FROM songs_by_user WHERE user_id = $1;

-- name: UpdateSong :exec
UPDATE songs SET title=$1, lyrics=$2, artist_id=$3 WHERE id=$4;

-- name: SearchSongsByUser :many
SELECT user_id, user_first_name, user_last_name, user_email, song_id, song_title, song_lyrics, 
    genre_id, genre, artist_id, artist_name FROM songs_by_user 
    WHERE user_id = $1 AND CONCAT(song_title, ' ', artist_name) ILIKE '%' || $2::text || '%' ORDER BY song_title LIMIT 10 OFFSET $3;

-- name: CreateSong :exec
INSERT INTO songs (title, lyrics, user_id, genre_id, artist_id) VALUES($1, $2, $3, $4, $5) RETURNING id;