-- +migrate Up
ALTER TABLE gems
ADD attachment BYTEA;

-- +migrate Down
ALTER TABLE gems
DROP COLUMN attachment;
