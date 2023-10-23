-- name: GetUser :one
SELECT * FROM "user" WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "user" (username, password, email, created_on, last_login, first_name, last_name) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;