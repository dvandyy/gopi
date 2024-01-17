CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "role" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "workspaces" (
  "id" bigserial PRIMARY KEY,
  "teams" bigint,
  "owner" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "boards" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "teams" (
  "id" bigserial PRIMARY KEY,
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

