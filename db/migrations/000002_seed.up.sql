-- Users
INSERT INTO users (first_name, last_name, email, password) 
VALUES 
    ('John', 'Doe', 'jdoe@example.com', 'password1'),
    ('Mary', 'Jane', 'mj@example.com', 'password2');

-- Composers
INSERT INTO composers (last_name, first_name)
VALUES
    ('Simon', 'Paul'),
    ('Lennon', 'John'),
    ('McCartney', 'Paul');

-- Performers
INSERT INTO performers (group_name, first_name, last_name)
VALUES
    ('Beatles', NULL, NULL),
    (NULL, 'Paul', 'Simon');

-- Songs
INSERT INTO songs (title, key, genre, capo, lyrics, user_id, composer_id, performer_id)
VALUES
    ('Lucy in the Sky with Diamonds', 'D', 'Rock', 3, '[A]Picture your[A7]self on a [D]boat on a [DM]river', 1),
    ('The Boxer', NULL, 'Folk', NULL, '[C]I am just a poor boy\nThough my story''s seldom [AM]told', 1);

