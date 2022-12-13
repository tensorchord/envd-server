table "image_info" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "owner_token" {
    null = false
    type = text
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
  primary_key {
    columns = [column.id]
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
  column "public_key" {
    null = false
    type = bytea
  }
  primary_key {
    columns = [column.id]
  }
  index "unique_login_name" {
    unique  = true
    columns = [column.login_name]
    type    = BTREE
  }
}
schema "public" {
}