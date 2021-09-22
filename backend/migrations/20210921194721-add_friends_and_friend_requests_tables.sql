-- +migrate Up
CREATE TABLE friends (
    first_user_id  BIGINT NOT NULL REFERENCES users (id),
    second_user_id BIGINT NOT NULL REFERENCES users (id),
    PRIMARY KEY (first_user_id, second_user_id),
    CHECK (first_user_id < second_user_id)
);

CREATE TABLE friend_requests (
    requester_id BIGINT NOT NULL REFERENCES users (id),
    receiver_id  BIGINT NOT NULL REFERENCES users (id),
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (requester_id, receiver_id)
);

-- +migrate Down
DROP TABLE friend_requests;
DROP TABLE friends;
