INSERT INTO users (first_name, last_name, email, password) 
VALUES 
    ('John', 'Doe', 'jdoe@example.com', 'password1'),
    ('Mary', 'Jane', 'mj@example.com', 'password2');

INSERT INTO composers (last_name, first_name)
VALUES
    ('Simon', 'Paul'),
    ('Lennon', 'John'),
    ('McCartney', 'Paul')

INSERT INTO performers (group_name, first_name, last_name)
VALUES
    ('Beatles', NULL, NULL),
    (NULL, 'Paul', 'Simon');

INSERT INTO songs (title, key, genre, capo, lyrics, user_id)
VALUES
    ('Lucy in the Sky with Diamonds', 'D', 'Rock', 3, '[A]Picture your[A7]self on a [D]boat on a [DM]river', 1),
    ('The Boxer', NULL, 'Folk', NULL, '[C]I am just a poor boy\nThough my story''s seldom [AM]told', 1);

    --     song_id SERIAL PRIMARY KEY,
    -- title VARCHAR(255) NOT NULL,
    -- "key" VARCHAR(5),
    -- genre genre NOT NULL,
    -- capo VARCHAR(20),
    -- lyrics TEXT NOT NULL,
    -- user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
