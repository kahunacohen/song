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

INSERT INTO songs (title, key, genre, capo, lyrics, user_id, artist_id)
VALUES
    ('Lucy in the Sky with Diamonds', 'D', 'Rock', 3, '[A]Picture your[A7]self on a [D]boat on a [DM]river', 1, 2),
    ('The Boxer', NULL, 'Folk', NULL, '[C]I am just a poor boy\nThough my story''s seldom [AM]told', 1, 1),
    ('Penny Lane', NULL, 'Rock', NULL, '[C]Penny Lane', 2, 2),
    ('Born in the USA', NULL, 'Rock', NULL, '[C]Born', 2, 3),
    ('This Land is Your Land', NULL, 'Folk', 2, '[C]This', 1, 4),
    ('Amazing Grace', 'C', 'Traditional', NULL, '[C]Amazing', 2, 5);
