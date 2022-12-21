-- +goose Up
-- modify "image_info" table
ALTER TABLE "public"."image_info" DROP COLUMN "owner_token", ADD COLUMN "login_name" text NOT NULL, ADD COLUMN "apt_packages" jsonb NOT NULL, ADD COLUMN "pypi_commands" jsonb NOT NULL, ADD COLUMN "services" jsonb NOT NULL;
-- create index "unique_digest" to table: "image_info"
CREATE UNIQUE INDEX "unique_digest" ON "public"."image_info" ("digest");

-- +goose Down
-- reverse: create index "unique_digest" to table: "image_info"
DROP INDEX "public"."unique_digest";
-- reverse: modify "image_info" table
ALTER TABLE "public"."image_info" DROP COLUMN "services", DROP COLUMN "pypi_commands", DROP COLUMN "apt_packages", DROP COLUMN "login_name", ADD COLUMN "owner_token" text NOT NULL;
