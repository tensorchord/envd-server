-- +goose Up
-- drop index "unique_digest" from table: "image_info"
DROP INDEX "public"."unique_digest";
-- create index "unique_login_name_and_digest" to table: "image_info"
CREATE UNIQUE INDEX "unique_login_name_and_digest" ON "public"."image_info" ("digest", "login_name");

-- +goose Down
-- reverse: create index "unique_login_name_and_digest" to table: "image_info"
DROP INDEX "public"."unique_login_name_and_digest";
-- reverse: drop index "unique_digest" from table: "image_info"
CREATE UNIQUE INDEX "unique_digest" ON "public"."image_info" ("digest");
