-- name: GetUser :one
SELECT * FROM users
WHERE identity_token = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (
  identity_token, public_key
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetImageInfo :one
SELECT * FROM image_info
WHERE owner_token = $1 AND name = $2 LIMIT 1;


-- name: ListImageByOwner :many
SELECT * FROM image_info
WHERE owner_token = $1;

-- name: CreateImageInfo :one
INSERT INTO image_info (
  owner_token, name, digest, created, size, labels
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;