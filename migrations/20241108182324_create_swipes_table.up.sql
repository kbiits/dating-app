CREATE TYPE "swipe_direction" AS ENUM (
  'like',
  'dislike'
);

CREATE TABLE "swipes" (
  "id" UUID UNIQUE PRIMARY KEY,
  "swiper_id" UUID NOT NULL,
  "swiped_id" UUID NOT NULL,
  "swipe_type" swipe_direction NOT NULL,
  "swipe_date" date NOT NULL DEFAULT CURRENT_DATE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON "swipes" ("swipe_date", "swiper_id");
CREATE INDEX ON "swipes" USING HASH ("swiper_id");

ALTER TABLE "swipes" ADD FOREIGN KEY ("swiper_id") REFERENCES "user_profiles" ("id");
ALTER TABLE "swipes" ADD FOREIGN KEY ("swiped_id") REFERENCES "user_profiles" ("id");