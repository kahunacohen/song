CREATE VIEW songs_by_user AS 
    SELECT users.id AS user_id, songs.id AS song_id, songs.title, songs.genre, songs.lyrics, artists.name AS artist_name 
    FROM songs 
    JOIN artists ON songs.artist_id = artists.id
    JOIN users ON songs.user_id = users.id;
