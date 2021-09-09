-- +migrate Up
CREATE TABLE gems (
    id      BIGSERIAL PRIMARY KEY,
    message VARCHAR
);

-- +migrate Down
DROP TABLE gems;
