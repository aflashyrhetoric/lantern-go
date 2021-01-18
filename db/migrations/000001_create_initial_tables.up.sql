CREATE TYPE "relationshiptype" AS ENUM (
  'spouse',
  'significant_other',
  'cousin',
  'sibling',
  'child',
  'parent',
  'enemy'
);
CREATE TABLE "people" (
  "id" int PRIMARY KEY,
  "first_name" text,
  "last_name" text,
  "career" text,
  "mobile" text,
  "email" int,
  "address" int,
  "dob" timestamp
);
CREATE TABLE "strengths" (
  "id" int,
  "person_id" int,
  "description"text 
);
CREATE TABLE "pressure_points" (
  "id" int,
  "person_id" int,
  "description"text 
);
CREATE TABLE "attendees" ("attendee_id" int, "event_id" int);
CREATE TABLE "notes" ("id" int, "person_id" int, "text" text);
CREATE TABLE "events" (
  "id" int PRIMARY KEY,
  "event_description" text,
  "notes" text,
  "date" timestamp
);
CREATE TABLE "relationship" (
  "person_one_id" int,
  "person_two_id" int,
  "relationship_type" relationshiptype
);
ALTER TABLE "strengths"
ADD FOREIGN KEY ("person_id") REFERENCES "people" ("id");
ALTER TABLE "pressure_points"
ADD FOREIGN KEY ("person_id") REFERENCES "people" ("id");
ALTER TABLE "attendees"
ADD FOREIGN KEY ("attendee_id") REFERENCES "people" ("id");
ALTER TABLE "attendees"
ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");
ALTER TABLE "notes"
ADD FOREIGN KEY ("person_id") REFERENCES "people" ("id");
ALTER TABLE "relationship"
ADD FOREIGN KEY ("person_one_id") REFERENCES "people" ("id");
ALTER TABLE "relationship"
ADD FOREIGN KEY ("person_two_id") REFERENCES "people" ("id");