
CREATE OR REPLACE FUNCTION update_timestamp_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  email VARCHAR NOT NULL UNIQUE,
  encrypted_password VARCHAR NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS hotels (
  id BIGSERIAL PRIMARY KEY,
  hotel_name VARCHAR NOT NULL,
  location VARCHAR NOT NULL,
  stars INT NOT NULL DEFAULT(0),
  created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);


CREATE TABLE IF NOT EXISTS room_types (
  id BIGSERIAL PRIMARY KEY,
  type VARCHAR NOT NULL UNIQUE,
  description VARCHAR NOT NULL,
  capacity INT NOT NULL DEFAULT(1),
  price_per_night DECIMAL NOT NULL CHECK(price_per_night > 0.0)
);

CREATE TABLE IF NOT EXISTS rooms (
  id BIGSERIAL PRIMARY KEY,
  room_number VARCHAR UNIQUE NOT NULL,
  room_type_id BIGINT NOT NULL REFERENCES room_types(id),
  hotel_id BIGINT NOT NULL REFERENCES hotels(id)
);



CREATE TRIGGER update_timestamp_users
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_column();


CREATE TRIGGER update_timestamp_hotels
BEFORE UPDATE ON hotels
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_column();


CREATE TRIGGER update_timestamp_rooms
BEFORE UPDATE ON rooms
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_column();


CREATE TRIGGER update_timestamp_rooms_types
BEFORE UPDATE ON room_types
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp_column();