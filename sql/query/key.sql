-- name: ListKeys :many
SELECT * FROM keys
WHERE login_name = $1;

-- name: GetKey :one
SELECT * FROM keys
WHERE login_name = $1 AND name = $2 LIMIT 1;

-- name: CreateKey :one
INSERT INTO keys (
  login_name, name, public_key
) VALUES (
  $1, $2, $3
)
RETURNING login_name, name, public_key;
