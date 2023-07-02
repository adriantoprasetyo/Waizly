CREATE TABLE "tickets" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL,
  "agent_id" int NULL,
  "status" bool NOT NULL DEFAULT TRUE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY (user_id) REFERENCES users (id)
);