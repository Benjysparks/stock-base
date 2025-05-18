-- +goose Up
CREATE TABLE stock (
    id  UUID PRIMARY KEY,
    stockname    TEXT UNIQUE NOT NULL,
    amount      INTEGER  NOT NULL,
    qty_type    TEXT NOT NULL,
    price_per   FLOAT NOT NULL
);

-- +goose Down
DROP TABLE stock;