CREATE VIEW artists_by_user AS 
    SELECT users.id AS user_id, artists.id AS artist_id, artists.name AS artist_name
    FROM artists 
    JOIN users ON artists.user_id = users.id;