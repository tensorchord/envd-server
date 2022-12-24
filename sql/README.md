# SQL

We use `goose` file format for migrations

## Generate golang code from SQL

We use [sqlc](https://github.com/kyleconroy/sqlc) for the SQL related codes. To generate the golang code:
- Install sqlc: `go install github.com/kyleconroy/sqlc/cmd/sqlc@latest`
- Execute `sqlc generate` at the project root

## Apply migration files

Run `atlas migrate apply --dir=file://schema?format=goose --url <dburl> --allow-dirty`. dburl is the database connection string for the database you want to change, such as `postgres://username:password@localhost:5432/database_name`.

## Write new migration tools

### Install atlas

Please find the guide at https://atlasgo.io/getting-started/

### Generate hcl file from current db schema

Run `atlas schema inspect --url "postgres://postgres:atlasdev@localhost:5432/postgres?sslmode=disable" --schema=public > atlas_schema.hcl`

### Rehash the atlas migration 

Run `atlas migrate hash --dir file://schema`

### Author a new migration

- `docker run --name some-postgres -e POSTGRES_PASSWORD=atlasdev -p5432:5432 -d postgres`
- Update the `atlas_schema.hcl` file
- Run `atlas migrate diff --dir file://schema/ --dir-format goose <summary> --to "file://atlas_schema.hcl" --dev-url "postgres://postgres:atlasdev@localhost:5432/postgres?sslmode=disable"`. This will generate a migration file from the exsiting schema files to the state described in `atlas_schema.hcl` file. The migration file will be created at `schema/<timestamp>_<summary>.sql`.
