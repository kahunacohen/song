CREATE VIEW songs_by_user AS 
    SELECT 
        users.id AS user_id,
        users.first_name AS user_first_name,
        users.last_name AS user_last_name,
        users.email AS user_email,
        songs.id AS song_id,
        songs.title AS song_title,
        songs.lyrics AS song_lyrics,
        genres.id AS genre_id,
        genres.name AS genre,
        artists.id AS artist_id,
        artists.name AS artist_name 
    FROM songs 
    JOIN artists ON songs.artist_id = artists.id
    JOIN genres ON songs.genre_id = genres.id
    JOIN users ON songs.user_id = users.id;
