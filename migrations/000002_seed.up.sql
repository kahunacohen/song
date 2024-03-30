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
    ('This Land is Your Land', NULL, 2, '[C]This', 1, 4, 2),
    ('Abc', NULL, 2, '[C]This', 1, 4, 2),
    ('Def', NULL, 2, '[C]This', 1, 4, 2),
    ('Fgh', NULL, 2, '[C]This', 1, 4, 2),
    ('Ijk', NULL, 2, '[C]This', 1, 4, 2),
    ('Lmn', NULL, 2, '[C]This', 1, 4, 2),
    ('Opq', NULL, 2, '[C]This', 1, 4, 2),
    ('Rst', NULL, 2, '[C]This', 1, 4, 2),
    ('Xyz', NULL, 2, '[C]This', 1, 4, 2),
    ('Drf', NULL, 2, '[C]This', 1, 4, 2),
    ('Mop', NULL, 2, '[C]This', 1, 4, 2),
    ('Penny Lane', NULL, NULL, '[C]Penny Lane', 2, 2, 1),
    ('Born in the USA', NULL, NULL, '[C]Born', 2, 3, 1),
    ('Amazing Grace', 'C', NULL, '[C]Amazing', 2, 5, 2);
