-- Create the users table
CREATE TABLE IF NOT EXISTS users
(
    id         serial PRIMARY KEY,
    name       TEXT        NOT NULL,
    password   TEXT        NOT NULL,
    email      TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP   NOT NULL default current_timestamp,
    updated_at TIMESTAMP   NOT NULL default current_timestamp
);

-- Create the assets table
CREATE TABLE IF NOT EXISTS assets
(
    asset_id       serial PRIMARY KEY,
    user_id        INT REFERENCES users (id),
    asset_name     VARCHAR(255),
    purchase_price NUMERIC,
    purchase_date  TIMESTAMP,
    current_price  NUMERIC,
    "current_date"   TIMESTAMP
);

-- Create the asset_price table
CREATE TABLE IF NOT EXISTS asset_price
(
    asset_price_id serial PRIMARY KEY,
    asset_id       INT REFERENCES assets (asset_id),
    price          NUMERIC,
    date           TIMESTAMP
    -- Add other fields as needed
);

-- Optionally, add foreign key constraints for relationships
ALTER TABLE assets
    ADD CONSTRAINT assets_users_fk
        FOREIGN KEY (user_id)
            REFERENCES users (id);

ALTER TABLE asset_price
    ADD CONSTRAINT asset_price_assets_fk
        FOREIGN KEY (asset_id)
            REFERENCES assets (asset_id);