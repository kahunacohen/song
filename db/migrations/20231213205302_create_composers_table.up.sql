CREATE TABLE composers (
    composer_id SERIAL PRIMARY KEY,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TYPE genre AS ENUM ('Rock', 'Pop', 'Hip-Hop', 'Country', 'Traditional');

CREATE TABLE songs (
    song_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    "key" VARCHAR(5),
    genre genre NOT NULL,
    capo VARCHAR(20),
    lyrics TEXT,
    composer_id INT REFERENCES composers(composer_id) ON DELETE CASCADE,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
