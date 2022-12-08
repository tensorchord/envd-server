-- name: GetUser :one
SELECT * FROM users
WHERE login_name = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (
  login_name, password_hash, public_key
) VALUES (
  $1, $2, $3
)
RETURNING login_name, public_key;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
