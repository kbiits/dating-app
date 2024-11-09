ALTER TABLE user_profiles 
    ADD COLUMN location VARCHAR(255),
    DROP COLUMN district_id;
