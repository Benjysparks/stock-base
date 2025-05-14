-- name: CreateStockItem :one
INSERT INTO stock (id, stockname, amount, qty_type, price_per)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;