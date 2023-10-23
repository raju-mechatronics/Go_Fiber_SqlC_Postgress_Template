-- create user table
CREATE TABLE "user" (
                        id SERIAL PRIMARY KEY,
                        username VARCHAR(255) UNIQUE NOT NULL,
                        password VARCHAR(255) NOT NULL,
                        email VARCHAR(255) UNIQUE NOT NULL,
                        created_on TIMESTAMP NOT NULL default now(),
                        last_login TIMESTAMP NOT NULL default now()
);