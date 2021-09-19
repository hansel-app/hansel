-- +migrate Up
CREATE TYPE gem_color AS ENUM ('purple', 'pink', 'blue', 'black', 'yellow', 'green');

ALTER TABLE gems
ADD color gem_color NOT NULL;


-- +migrate Down
ALTER TABLE gems
DROP COLUMN color;

DROP TYPE gem_color;
