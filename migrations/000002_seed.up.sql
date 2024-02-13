INSERT INTO users (first_name, last_name, email, password) 
VALUES 
    ('John', 'Doe', 'jdoe@example.com', 'password1'),
    ('Mary', 'Jane', 'mj@example.com', 'password2');

INSERT INTO composers (last_name, first_name, name)
VALUES
    ('Simon', 'Paul', NULL),
    ('Lennon', 'John', NULL),
    ('McCartney', 'Paul', NULL),
    (NULL, NULL, 'Lennon–McCartney'),
    ('Springsteen', 'Bruce', NULL),
    ('Guthrey', 'Woody', NULL);

INSERT INTO performers (name, first_name, last_name)
VALUES
    ('Beatles', NULL, NULL),
    (NULL, 'Paul', 'Simon');

INSERT INTO songs (title, key, genre, capo, lyrics, user_id, composer_id, performer_id)
VALUES
    ('Lucy in the Sky with Diamonds', 'D', 'Rock', 3, '[A]Picture your[A7]self on a [D]boat on a [DM]river', 1, 4, 1),
    ('The Boxer', NULL, 'Folk', NULL, '[C]I am just a poor boy\nThough my story''s seldom [AM]told', 1, 1, 2),
    ('Penny Lane', NULL, 'Rock', NULL, '[C]Penny Lane', 2, 4, 1),
    ('Born in the USA', NULL, 'Rock', NULL, '[C]Born', 2, 5, NULL),
    ('This Land is Your Land', NULL, 'Folk', 2, '[C]This', 1, 6, NULL),
    ('Amazing Grace', 'C', 'Traditional', NULL, '[C]Amazing', 2, NULL, NULL);
