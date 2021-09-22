-- +migrate Up
ALTER TABLE users
    DROP COLUMN email;

-- +migrate Down
ALTER TABLE users
    ADD email VARCHAR UNIQUE NOT NULL;
