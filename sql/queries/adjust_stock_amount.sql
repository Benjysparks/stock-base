-- name: AdjustStockAmount :exec

UPDATE stock
SET amount = amount + $1
WHERE stockname = $2;