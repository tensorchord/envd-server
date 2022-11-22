-- Users
CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  identity_token text NOT NULL,
  public_key bytea NOT NULL
);


-- Image info
CREATE TABLE IF NOT EXISTS image_info (
  id BIGSERIAL PRIMARY KEY,
  owner_token text NOT NULL,
  name text NOT NULL,
  digest text NOT NULL,
  created bigint NOT NULL,
  size bigint NOT NULL,
  labels JSONB
);