CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE options AS ENUM('Flood', 'Landslide', 'Fire', 'Construction', 'Accident');

CREATE TABLE sensors (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT,
    latitude FLOAT8,
    longitude FLOAT8
);

INSERT INTO sensors (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);
INSERT INTO sensors (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);
INSERT INTO sensors (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);
INSERT INTO sensors (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);
INSERT INTO sensors (name, latitude, longitude) VALUES ('Sensor Name 1', -23.5718, -46.708);

CREATE TABLE sensors_log (
    id SERIAL PRIMARY KEY,
    sensor_id UUID REFERENCES sensors(id),
    data TEXT,
    timestamp TIMESTAMP
);

-- CREATE TABLE air_quality (
--     id SERIAL PRIMARY KEY,
--     sensor_id UUID REFERENCES sensors(id),
--     co2 FLOAT8,
--     co FLOAT8,
--     no2 FLOAT8,
--     mp10 FLOAT8,
--     mp25 FLOAT8,
--     timestamp TIMESTAMP
-- )

-- CREATE TABLE radiation (
--     id SERIAL PRIMARY KEY,
--     sensor_id UUID REFERENCES sensors(id),
--     rad FLOAT8,
--     timestamp TIMESTAMP
-- )

CREATE TABLE alerts (
    id SERIAL PRIMARY KEY,
    latitude FLOAT8,
    longitude FLOAT8,
    option OPTIONS,
    timestamp TIMESTAMP
);