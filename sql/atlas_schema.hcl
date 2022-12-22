table "image_info" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "name" {
    null = false
    type = text
  }
  column "digest" {
    null = false
    type = text
  }
  column "created" {
    null = false
    type = bigint
  }
  column "size" {
    null = false
    type = bigint
  }
  column "labels" {
    null = true
    type = jsonb
  }
  column "login_name" {
    null = false
    type = character_varying(100)
  }
  column "apt_packages" {
    null = false
    type = jsonb
  }
  column "pypi_commands" {
    null = false
    type = jsonb
  }
  column "services" {
    null = false
    type = jsonb
  }
  primary_key {
    columns = [column.id]
  }
  index "unique_digest" {
    unique  = true
    columns = [column.digest]
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "login_name" {
    null = false
    type = character_varying(100)
  }
  column "password_hash" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
  index "unique_login_name" {
    unique  = true
    columns = [column.login_name]
  }
}
table "keys" {
  schema = schema.public
    column "id" {
    null = false
    type = bigserial
  }
  column "name" {
    null = false
    type = text
  }
  column "login_name" {
    null = false
    type = character_varying(100)
  }
  column "public_key" {
    null = false
    type = bytea
  }
  primary_key {
    columns = [column.id]
  }
  index "unique_login_name_and_key" {
    unique  = true
    columns = [column.login_name, column.name]
  }
}
schema "public" {
}
