-- +goose Up
-- Create "urls" table
CREATE TABLE "public"."urls" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "original_url" text NOT NULL,
  "short_slug" character varying(10) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);

-- Create index "short_slug_unique" to table: "urls"
CREATE UNIQUE INDEX "short_slug_unique" ON "public"."urls" ("short_slug");

-- +goose Down
-- Drop the table and index
DROP INDEX IF EXISTS "short_slug_unique";
DROP TABLE IF EXISTS "public"."urls";