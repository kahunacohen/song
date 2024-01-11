
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
    last_name VARCHAR(255),
    first_name VARCHAR(255),
    "name" VARCHAR(255),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    CONSTRAINT check_not_null_columns
    CHECK ((last_name IS NOT NULL AND first_name IS NOT NULL) OR name IS NOT NULL)
);
CREATE TABLE performers (
    id SERIAL PRIMARY KEY,
    last_name VARCHAR(255),
    first_name VARCHAR(255),
    "name" VARCHAR(255),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    CONSTRAINT check_not_null_columns
    CHECK ((last_name IS NOT NULL AND first_name IS NOT NULL) OR name IS NOT NULL)
);
CREATE TYPE genre AS ENUM ('Bluegrass', 'Blues', 'Rock', 'Pop', 'Hip-Hop', 'Country', 'Traditional', 'Folk');

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    public_domain BOOLEAN NOT NULL DEFAULT FALSE,
    "key" VARCHAR(5),
    genre genre NOT NULL,
    capo SMALLINT,
    lyrics TEXT NOT NULL,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    composer_id INT REFERENCES composers(id),
    performer_id INT REFERENCES performers(id),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
