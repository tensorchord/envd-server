-- +goose Up
CREATE TABLE IF NOT EXISTS public.users (
  id BIGSERIAL PRIMARY KEY,
  identity_token text NOT NULL,
  public_key bytea NOT NULL
);

-- +goose Down
DROP TABLE users