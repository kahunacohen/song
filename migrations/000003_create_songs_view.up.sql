CREATE VIEW songs_by_user AS 
    SELECT songs.id, songs.title, songs.genre, songs.lyrics, artists.name 
    FROM songs 
    JOIN artists ON songs.artist_id = artists.id
    JOIN users ON songs.user_id = users.id;
