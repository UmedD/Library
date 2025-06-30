CREATE TABLE IF NOT EXISTS books
(
    id      SERIAL      PRIMARY KEY,
    name    TEXT        NOT NULL,
    title   TEXT        NOT NULL,
    authors INT[]       NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS authors
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS book_authors
(
    book_id   INTEGER NOT NULL REFERENCES books (id) ON DELETE CASCADE,
    author_id INTEGER NOT NULL REFERENCES authors (id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, author_id)
);

drop table books;

drop table users;

CREATE TABLE IF NOT EXISTS users
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE  NOT NULL,
    email    VARCHAR(100) UNIQUE NOT NULL,
    password TEXT                NOT NULL,
    role     VARCHAR(10)
);

ALTER TABLE users
    ALTER COLUMN role SET DEFAULT 'user';

ALTER TABLE books
    ADD COLUMN author_id INTEGER NULL
        REFERENCES authors(id) ON DELETE RESTRICT;

SELECT
    b.id,
    b.name,
    b.title,
    b.author_id,
    a.name AS author_name
FROM books b
         JOIN authors a ON a.id = b.author_id
WHERE b.id = 3;
