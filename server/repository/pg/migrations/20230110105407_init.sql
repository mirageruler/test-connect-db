-- +goose Up
-- +goose StatementBegin
-- uuid-ossp is a contrib module, so it isn't loaded into the server by default. You must load it into your database to use it.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE "users" (
  "id" uuid,
  "age" bigint,
  "name" varchar(42),
  "is_admin" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS "users"
-- +goose StatementEnd
