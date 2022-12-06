-- Users
CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  login_name varchar(100) NOT NULL,
  password_hash text NOT NULL,
  public_key bytea NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS users_login_name 
ON users (login_name);

ALTER TABLE users 
DROP CONSTRAINT IF EXISTS unique_login_name;

ALTER TABLE users 
ADD CONSTRAINT unique_login_name 
UNIQUE USING INDEX users_login_name;

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
