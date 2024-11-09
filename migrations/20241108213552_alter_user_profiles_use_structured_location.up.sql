ALTER TABLE user_profiles
    DROP COLUMN location,
    ADD COLUMN district_id CHAR(6);

ALTER TABLE "user_profiles"
ADD CONSTRAINT fk_profiles_district_id FOREIGN KEY ("district_id") REFERENCES "districts" ("id");

CREATE INDEX idx_user_profiles_district_id ON user_profiles USING HASH (district_id);