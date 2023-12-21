
CREATE TABLE  users (
    user_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
CREATE TABLE composers (
    composer_id SERIAL PRIMARY KEY,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
CREATE TABLE performers (
    performer_id SERIAL PRIMARY KEY,
    last_name VARCHAR(255),
    first_name VARCHAR(255),
    group_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
CREATE TYPE genre AS ENUM ('Bluegrass', 'Blues', 'Rock', 'Pop', 'Hip-Hop', 'Country', 'Traditional', 'Folk');

CREATE TABLE songs (
    song_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    "key" VARCHAR(5),
    genre genre NOT NULL,
    capo SMALLINT,
    lyrics TEXT NOT NULL,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE composer_song (
    composer_id INT REFERENCES composers(composer_id),
    song_id INT REFERENCES songs(song_id),
    PRIMARY KEY (composer_id, song_id)
);

CREATE TABLE performer_song (
    performer_id INT REFERENCES performers(performer_id),
    song_id INT REFERENCES songs(song_id),
    PRIMARY KEY (performer_id, song_id)
);
