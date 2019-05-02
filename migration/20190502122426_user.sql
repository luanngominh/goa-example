-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."users" (
    "id" uuid NOT NULL,
    "email" TEXT NOT NULL,
    "fullname" TEXT NOT NULL,
    "password_digest" TEXT NOT NULL,
    "verify_code" CHAR(6) DEFAULT '',
    "code_expire_time" TIMESTAMPTZ DEFAULT NULL,
    "created_at" TIMESTAMPTZ DEFAULT now(),
    "deleted_at" TIMESTAMPTZ DEFAULT NULL,

    CONSTRAINT "user_pk" PRIMARY KEY ("id")
)

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA "public" CASCADE