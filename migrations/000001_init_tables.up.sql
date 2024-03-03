
CREATE TABLE  users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
CREATE TABLE artists (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    public_domain BOOLEAN NOT NULL DEFAULT FALSE,
    "key" VARCHAR(5),
    capo SMALLINT,
    lyrics TEXT NOT NULL,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    genre_id INT REFERENCES genres(id) ON DELETE CASCADE NOT NULL,
    artist_id INT REFERENCES artists(id),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);