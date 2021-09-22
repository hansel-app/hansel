-- +migrate Up
ALTER TABLE users
    ADD display_name VARCHAR NOT NULL DEFAULT '',
    ADD avatar       BYTEA;

-- +migrate Down
ALTER TABLE users
    DROP COLUMN display_name,
    DROP COLUMN avatar;
