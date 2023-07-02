CREATE TABLE "comments" (
  "id" serial PRIMARY KEY,
  "ticket_id" int NOT NULL,
  "text" varchar NULL,
  "user_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (ticket_id) REFERENCES tickets (id)

);