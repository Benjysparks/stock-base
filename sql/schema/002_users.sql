-- +goose Up
CREATE TABLE users (
    id  UUID PRIMARY KEY,
    user_name    TEXT UNIQUE NOT NULL,
    pass_word      TEXT  NOT NULL
);

-- +goose Down
DROP TABLE history;