env "local" {
  url = "postgres://lars:1602@localhost:5432/postgres?sslmode=disable"
  dev = "postgres://lars:1602@localhost:5432/postgres?sslmode=disable"
  migration {
    dir = "file://internal/db/migrations"
  }
}
