-- name: GetArtistsByUser :many
SELECT * FROM artists_by_user WHERE user_id = $1 ORDER BY artist_name LIMIT 10 OFFSET $2;

-- name: GetArtist :one
SELECT name FROM artists WHERE id = $1;

-- name: UpdateArtist :exec
UPDATE artists SET name=$1 WHERE id=$2;