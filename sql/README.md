# SQL

We use `goose` file format for migrations

## Generate golang code from SQL

We use [sqlc](https://github.com/kyleconroy/sqlc) for the SQL related codes. To generate the golang code:
- Install sqlc: `go install github.com/kyleconroy/sqlc/cmd/sqlc@latest`
- Execute `sqlc generate` at the project root

## Apply migration files

Run `atlas migrate apply --dir=file://schema --url <dburl>`. dburl is the database connection string for the database you want to change, such as `postgres://username:password@localhost:5432/database_name`.

## Write new migration tools

### Install atlas

Please find the guide at https://atlasgo.io/getting-started/

### Generate hcl file from current db schema

Run `atlas schema inspect --url "postgres://postgres:atlasdev@localhost:5432/postgres?sslmode=disable" > atlas_schema.hcl`

### Author a new migration

Option 1: Manually write the migration file

Run `atlas migrate new --dir file://schema/ --dir-format goose <summary>`, summary will be append to the migration filename. The migration file will be created at `schema/<timestamp>_<summary>.sql`.

Option 2: Automatically generate migration file based on current db schema

1. You need to spin up an empty db as the `dev database` for atlas to compare with the current db schema. You can use the following command to spin up a docker container for the `dev database`:
```
docker run --name atlasdev -e POSTGRES_PASSWORD=atlasdev -p 5432:5432 -d postgres
```

2. Generate hcl file following the guide above

Run `atlas schema inspect --url "postgres://postgres:atlasdev@localhost:5432/postgres?sslmode=disable" > atlas_schema.hcl`. Or you can manually modify the `atlas_schema.hcl`


3. Run `atlas migrate diff --dir file://schema/ --dir-format goose <summary> --to "file://atlas_schema.hcl" --dev-url "postgres://postgres:atlasdev@localhost:5432/postgres?sslmode=disable"`. This will generate a migration file from the exsiting schema files to the state described in `atlas_schema.hcl` file. The migration file will be created at `schema/<timestamp>_<summary>.sql`.