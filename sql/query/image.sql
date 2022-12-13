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
