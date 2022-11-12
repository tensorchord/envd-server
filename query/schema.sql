-- Users
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  identity_token text NOT NULL,
  public_key bytea NOT NULL
);


-- Image info
CREATE TABLE image_info (
  id BIGSERIAL PRIMARY KEY,
  owner_token text NOT NULL,
  name text NOT NULL,
  digest text NOT NULL,
  created bigint NOT NULL,
  size bigint NOT NULL,
  labels JSONB
);