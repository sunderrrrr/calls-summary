CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       pass_hash VARCHAR(255) NOT NULL,
                       user_type SMALLINT NOT NULL DEFAULT 0 CHECK (user_type BETWEEN 0 AND 2)

);