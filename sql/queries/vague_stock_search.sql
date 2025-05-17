-- name: VagueStockSearch :many

SELECT * FROM stock 
WHERE stockname LIKE $1;