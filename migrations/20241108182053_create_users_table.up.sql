CREATE TYPE "profile_status" AS ENUM (
  'unverified',
  'verified',
  'deactivated'
);

CREATE TYPE "gender" AS ENUM (
  'male',
  'female'
);

CREATE TABLE "users" (
  "id" UUID UNIQUE PRIMARY KEY,
  "email" VARCHAR UNIQUE NOT NULL,
  "password" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "user_profiles" (
  "id" UUID UNIQUE PRIMARY KEY,
  "user_id" UUID,
  "display_name" VARCHAR NOT NULL,
  "bio" TEXT,
  "gender" gender,
  "date_of_birth" TIMESTAMP,
  "location" VARCHAR,
  "profile_pic_url" VARCHAR,
  "status" profile_status NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON "user_profiles" USING HASH ("user_id");
ALTER TABLE "user_profiles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");