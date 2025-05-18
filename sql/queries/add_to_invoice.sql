-- name: AddToInvoice :one
INSERT INTO invoice (id, user_name, stockname, amount, qty_type, price_per, total_price, created_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    NOW()
)
RETURNING *;