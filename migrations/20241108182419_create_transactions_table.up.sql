
CREATE TABLE "premium_packages" (
  "id" UUID UNIQUE PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "description" TEXT,
  "price" DECIMAL(10,2) NOT NULL,
  "validity" INT, -- validity in minutes
  "config" JSONB, -- This column will store the package configuration, such as quota per day/month/etc
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "user_purchases" (
  "id" UUID UNIQUE PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "package_id" UUID NOT NULL,
  "purchase_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_active" BOOLEAN NOT NULL DEFAULT false
);

CREATE INDEX ON "user_purchases" ("is_active", "user_id");

ALTER TABLE "user_purchases" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_purchases" ADD FOREIGN KEY ("package_id") REFERENCES "premium_packages" ("id");

insert into premium_packages (id, name, description, price, validity, config)
values (
	UUID_GENERATE_V4(), 'Unlimited Quota', 'Unlimited quota package', '50000'::DECIMAL, 300, '{
  "unlimited": true
}'::JSONB
);