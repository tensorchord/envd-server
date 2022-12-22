-- +goose Up
-- modify "image_info" table
ALTER TABLE "public"."image_info" ALTER COLUMN "login_name" TYPE character varying(100);
-- modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "public_key";
-- create "keys" table
CREATE TABLE "public"."keys" ("id" bigserial NOT NULL, "name" text NOT NULL, "login_name" character varying(100) NOT NULL, "public_key" bytea NOT NULL, PRIMARY KEY ("id"));
-- create index "unique_login_name_and_key" to table: "keys"
CREATE UNIQUE INDEX "unique_login_name_and_key" ON "public"."keys" ("login_name", "name");

-- +goose Down
-- reverse: create index "unique_login_name_and_key" to table: "keys"
DROP INDEX "public"."unique_login_name_and_key";
-- reverse: create "keys" table
DROP TABLE "public"."keys";
-- reverse: modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "public_key" bytea NOT NULL;
-- reverse: modify "image_info" table
ALTER TABLE "public"."image_info" ALTER COLUMN "login_name" TYPE text;
