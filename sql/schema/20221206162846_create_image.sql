-- +goose Up
CREATE TABLE IF NOT EXISTS image_info (
  id BIGSERIAL PRIMARY KEY,
  owner_token text NOT NULL,
  name text NOT NULL,
  digest text NOT NULL,
  created bigint NOT NULL,
  size bigint NOT NULL,
  labels JSONB
);

-- +goose Down
DROP TABLE image_info