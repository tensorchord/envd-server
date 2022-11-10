CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  identity_token text NOT NULL,
  public_key bytea NOT NULL
);