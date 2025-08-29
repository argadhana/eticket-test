-- users table
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR NOT NULL,
                       password_hash VARCHAR NOT NULL,
                       role VARCHAR NOT NULL,
                       fullname VARCHAR NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- stations table
CREATE TABLE stations (
                          id SERIAL PRIMARY KEY,
                          station_id VARCHAR UNIQUE NOT NULL, -- e.g. "ST01"
                          name VARCHAR NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- devices table
CREATE TABLE devices (
                         id SERIAL PRIMARY KEY,
                         device_id VARCHAR UNIQUE NOT NULL,  -- e.g. "GATE_A1"
                         station_id INTEGER NOT NULL REFERENCES stations(id),
                         last_seen TIMESTAMP,
                         firmware_version VARCHAR,
                         config_version VARCHAR,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- cards table
CREATE TABLE cards (
                       id SERIAL PRIMARY KEY,
                       card_id VARCHAR UNIQUE NOT NULL,  -- e.g. "CARD101"
                       balance INTEGER DEFAULT 0,
                       status VARCHAR CHECK (status IN ('ACTIVE', 'BLOCKED')),
                       txn_counter INTEGER DEFAULT 0,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- fares table
CREATE TABLE fares (
                       id SERIAL PRIMARY KEY,
                       fare_version VARCHAR NOT NULL,
                       from_station VARCHAR NOT NULL REFERENCES stations(station_id),
                       to_station VARCHAR NOT NULL REFERENCES stations(station_id),
                       amount INTEGER NOT NULL,
                       effective_from TIMESTAMP NOT NULL,
                       effective_to TIMESTAMP
);

-- trips table
CREATE TABLE trips (
                       id SERIAL PRIMARY KEY,
                       trip_id VARCHAR UNIQUE NOT NULL,  -- e.g. "TRX1001"
                       card_id VARCHAR NOT NULL REFERENCES cards(card_id),
                       entry_station VARCHAR NOT NULL REFERENCES stations(station_id),
                       exit_station VARCHAR REFERENCES stations(station_id),
                       entry_time TIMESTAMP NOT NULL,
                       exit_time TIMESTAMP,
                       amount INTEGER,
                       status VARCHAR CHECK (status IN ('IN_PROGRESS', 'COMPLETED', 'SYNC_PENDING')),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- transactions table
CREATE TABLE transactions (
                              id SERIAL PRIMARY KEY,
                              trip_id VARCHAR NOT NULL REFERENCES trips(trip_id),
                              local_txn_id VARCHAR NOT NULL,
                              device_id VARCHAR NOT NULL REFERENCES devices(device_id),
                              card_id VARCHAR NOT NULL REFERENCES cards(card_id),
                              txn_counter INTEGER NOT NULL,
                              amount INTEGER NOT NULL,
                              before_balance INTEGER NOT NULL,
                              after_balance INTEGER NOT NULL,
                              timestamp TIMESTAMP NOT NULL,
                              proof VARCHAR,
                              status VARCHAR CHECK (status IN ('RECEIVED', 'PROCESSED', 'REJECTED')),
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- sync_logs table
CREATE TABLE sync_logs (
                           id SERIAL PRIMARY KEY,
                           device_id VARCHAR NOT NULL REFERENCES devices(device_id),
                           batch_id VARCHAR NOT NULL,
                           synced_count INTEGER DEFAULT 0,
                           failed JSON,
                           raw_payload JSON,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
