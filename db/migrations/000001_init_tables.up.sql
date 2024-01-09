
CREATE TABLE  users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
CREATE TABLE composers (
    id SERIAL PRIMARY KEY,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
CREATE TABLE performers (
    id SERIAL PRIMARY KEY,
    last_name VARCHAR(255),
    first_name VARCHAR(255),
    group_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
CREATE TYPE genre AS ENUM ('Bluegrass', 'Blues', 'Rock', 'Pop', 'Hip-Hop', 'Country', 'Traditional', 'Folk');

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    "key" VARCHAR(5),
    genre genre NOT NULL,
    capo SMALLINT,
    lyrics TEXT NOT NULL,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    composer_id INT REFERENCES composers(id)
    performer_id INT REFERENCES performers(id)
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
