-- +migrate Up
CREATE TABLE users (
    id       BIGSERIAL PRIMARY KEY,
    username VARCHAR UNIQUE
);

CREATE TABLE gems (
    id          BIGSERIAL PRIMARY KEY,
    message     VARCHAR,
    latitude    DECIMAL(11, 8) NOT NULL,
    longitude   DECIMAL(11, 8) NOT NULl,
    creator_id  BIGSERIAL NOT NULL REFERENCES users(id),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    receiver_id BIGSERIAL NOT NULL REFERENCES users(id),
    received_at TIMESTAMP
);

-- +migrate Down
DROP TABLE users;
DROP TABLE gems;
