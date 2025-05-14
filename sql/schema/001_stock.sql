-- +goose Up
CREATE TABLE stock (
    id  UUID PRIMARY KEY,
    stockname    TEXT NOT NULL,
    amount      INT  NOT NULL,
    qty_type    TEXT NOT NULL,
    price_per   FLOAT NOT NULL
);

-- +goose Down
DROP TABLE stock;