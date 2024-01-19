CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "unique_id" uuid NOT NULL,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "role" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "workspaces" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "unique_id" uuid NOT NULL,
  "teams" bigint,
  "owner" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "boards" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "unique_id" uuid NOT NULL,
  "title" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "teams" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "unique_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "boards" bigint,
  "members" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("email");

ALTER TABLE "workspaces" ADD FOREIGN KEY ("teams") REFERENCES "teams" ("id");

ALTER TABLE "workspaces" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "teams" ADD FOREIGN KEY ("boards") REFERENCES "boards" ("id");

ALTER TABLE "teams" ADD FOREIGN KEY ("members") REFERENCES "users" ("id");
