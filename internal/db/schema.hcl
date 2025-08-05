schema "public" {
  comment = "Default schema"
}

table "urls" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
    default = sql("gen_random_uuid()")
  }
  column "original_url" {
    null = false
    type = text
  }
  column "short_slug" {
    null = false
    type = varchar(10)
  }
  column "created_at" {
    null = false
    type = timestamp
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
  index "short_slug_unique" {
    unique = true
    columns = [column.short_slug]
  }
}
