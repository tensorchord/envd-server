-- +goose Up
-- modify "users" table
-- atlas:nolint
ALTER TABLE "public"."users" DROP COLUMN "identity_token", ADD COLUMN "login_name" character varying(100) NOT NULL, ADD COLUMN "password_hash" text NOT NULL;
-- create index "unique_login_name" to table: "users"
-- atlas:nolint
CREATE UNIQUE INDEX "unique_login_name" ON "public"."users" ("login_name");

-- +goose Down
-- reverse: create index "unique_login_name" to table: "users"
-- atlas:nolint
DROP INDEX "public"."unique_login_name";
-- reverse: modify "users" table
-- atlas:nolint
ALTER TABLE "public"."users" DROP COLUMN "password_hash", DROP COLUMN "login_name", ADD COLUMN "identity_token" text NOT NULL;
