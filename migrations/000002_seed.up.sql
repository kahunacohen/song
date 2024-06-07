INSERT INTO users (first_name, last_name, email, password) 
VALUES 
    ('John', 'Doe', 'jdoe@example.com', 'password1'),
    ('Mary', 'Jane', 'mj@example.com', 'password2');

INSERT INTO artists (name, user_id)
VALUES
    ('Simon and Garfunkel', 1),
    ('Beatles', 1),
    ('Bruce Springsteen', 2),
    ('Woody Gurthrie', 1),
    ('Traditional', 2);

INSERT INTO genres (name)
VALUES
    ('Rock'),
    ('Folk'),
    ('Bluegrass');

INSERT INTO songs (title, key, capo, lyrics, user_id, artist_id, genre_id)
VALUES
    ('Lucy in the Sky with Diamonds', 'D', 3, '[A]Picture your[A7]self on a [D]boat on a [DM]river', 1, 2, 1),
    ('The Boxer', NULL, NULL, '[C]I am just a poor boy\nThough my story''s seldom [AM]told', 1, 1, 2),
    ('This Land is Your Land', NULL, 2, '[C]This', 1, 4, 2);
