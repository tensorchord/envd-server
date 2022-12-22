-- name: GetImageInfoByName :one
SELECT * FROM image_info
WHERE login_name = $1 AND name = $2 LIMIT 1;

-- name: GetImageInfoByDigest :one
SELECT * FROM image_info
WHERE login_name = $1 AND digest = $2 LIMIT 1;

-- name: ListImageByOwner :many
SELECT * FROM image_info
WHERE login_name = $1;

-- name: CreateImageInfo :one
INSERT INTO image_info (
  login_name, name, digest, created, size, 
  labels, apt_packages, pypi_commands, services
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;
