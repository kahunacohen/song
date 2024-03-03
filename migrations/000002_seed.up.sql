INSERT INTO users (first_name, last_name, email, password) 
VALUES 
    ('John', 'Doe', 'jdoe@example.com', 'password1'),
    ('Mary', 'Jane', 'mj@example.com', 'password2');

INSERT INTO artists (name)
VALUES
    ('Simon and Garfunkel'),
    ('Beatles'),
    ('Bruce Springsteen'),
    ('Woody Gurthrie'),
    ('Traditional');

INSERT INTO genres (name)
VALUES
    ('Rock'),
    ('Folk'),
    ('Bluegrass');

INSERT INTO songs (title, key, capo, lyrics, user_id, artist_id, genre_id)
VALUES
    ('Lucy in the Sky with Diamonds', 'D', 3, '[A]Picture your[A7]self on a [D]boat on a [DM]river', 1, 2, 1),
    ('The Boxer', NULL, NULL, '[C]I am just a poor boy\nThough my story''s seldom [AM]told', 1, 1, 2),
    ('Penny Lane', NULL, NULL, '[C]Penny Lane', 2, 2, 1),
    ('Born in the USA', NULL, NULL, '[C]Born', 2, 3, 1),
    ('This Land is Your Land', NULL, 2, '[C]This', 1, 4, 2),
    ('Amazing Grace', 'C', NULL, '[C]Amazing', 2, 5, 2);
