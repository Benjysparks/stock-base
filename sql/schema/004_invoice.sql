-- +goose Up
CREATE TABLE invoice (
    id  UUID PRIMARY KEY,
    user_name    TEXT  NOT NULL,
    stockname    TEXT UNIQUE NOT NULL,
    amount      INTEGER  NOT NULL,
    qty_type    TEXT NOT NULL,
    price_per   FLOAT NOT NULL,
    total_price FLOAT NOT NULL,
    created_at  TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE invoice;