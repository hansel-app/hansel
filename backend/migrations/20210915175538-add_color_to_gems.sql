-- +migrate Up
ALTER TABLE gems
ADD color bigint NOT NULL;


-- +migrate Down
ALTER TABLE gems
DROP COLUMN color;
