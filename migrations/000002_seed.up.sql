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
    ('Traditional', 2),
    ('Eagles', 2),
    ('Bob Dylan', 1);

INSERT INTO genres (name)
VALUES
    ('Rock'),
    ('Folk'),
    ('Bluegrass');

INSERT INTO songs (title, key, capo, lyrics, user_id, artist_id, genre_id)
VALUES
    ('Lucy in the Sky with Diamonds', 'D', 3, '[A]Picture your[A7]self on a [D]boat on a [DM]river', 1, 2, 1),
    ('The Boxer', NULL, NULL, '[C]I am just a poor boy\nThough my story''s seldom [AM]told', 1, 1, 2),
    ('This Land is Your Land', NULL, 2, '[C]This', 1, 4, 2),
    ('Take it Easy', NULL, 2, '[C]This', 2, 6, 1),
    ('Like a Rolling Stone', NULL, 2, '[C]This', 1, 7, 1),
    ('Hey Jude', NULL, 2, '[C]This', 1, 2, 1),
    ('Hard Days Night', NULL, 2, '[C]This', 1, 2, 1),
    ('Come Together', NULL, 2, '[C]This', 1, 2, 1),
    ('Yellow Submarine', NULL, 2, '[C]This', 1, 2, 1),
    ('Should Have Known Better', NULL, 2, '[C]This', 1, 2, 1),
    ('I Wanna Hold Your Hand', NULL, 2, '[C]This', 1, 2, 1),
    ('Yesterday', NULL, 2, '[C]This', 1, 2, 1);


