-- +migrate Up
ALTER TABLE gems
ADD color BIGINT NOT NULL;

-- +migrate Down
ALTER TABLE gems
DROP COLUMN color;
