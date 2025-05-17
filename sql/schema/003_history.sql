-- +goose Up
CREATE TABLE history (
    id  UUID PRIMARY KEY,
    user_name    TEXT  NOT NULL,
    stockname      TEXT  NOT NULL,
    edit_time   TIMESTAMP NOT NULL,
    adjustment  TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;