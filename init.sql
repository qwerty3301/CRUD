\c postgres;

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50),
                       email VARCHAR(100) UNIQUE,
                       password VARCHAR(100)
);