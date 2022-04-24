CREATE TABLE "users" if not exists (
  "id" SERIAL PRIMARY KEY,
  "email" text UNIQUE,
  "password" text NOT NULL,
  "created_at" timestamp
);

CREATE TABLE "people" if not exists (
  "id" SERIAL PRIMARY KEY,
  "first_name" text,
  "last_name" text,
  "career" text,
  "mobile" text,
  "email" text,
  "address" text,
  "dob" timestamp,
  "user_id" int
);

CREATE TABLE "pressure_points" if not exists (
  "id" SERIAL PRIMARY KEY,
  "person_id" int,
  "description" text
);

CREATE TABLE "notes" if not exists (
  "id" SERIAL PRIMARY KEY,
  "person_id" int,
  "text" text
);

CREATE TABLE "relationship" if not exists (
  "id" SERIAL PRIMARY KEY,
  "person_one_id" int,
  "person_two_id" int,
  "relationship_type" text
);

ALTER TABLE "people" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "pressure_points" ADD FOREIGN KEY ("person_id") REFERENCES "people" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("person_id") REFERENCES "people" ("id");

ALTER TABLE "relationship" ADD FOREIGN KEY ("person_one_id") REFERENCES "people" ("id");

ALTER TABLE "relationship" ADD FOREIGN KEY ("person_two_id") REFERENCES "people" ("id");
