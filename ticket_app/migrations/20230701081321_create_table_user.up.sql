CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NULL,
  "role" varchar NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);