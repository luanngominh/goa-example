-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."notes" (
    "id" uuid NOT NULL,
    "data" TEXT,
    "user_id" uuid NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT now(),
    "deleted_at" TIMESTAMPTZ DEFAULT NULL,

    CONSTRAINT "note_pk" PRIMARY KEY ("id"),
    CONSTRAINT "usr_note_fk" FOREIGN KEY ("user_id") REFERENCES users("id")
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA "public" CASCADE;