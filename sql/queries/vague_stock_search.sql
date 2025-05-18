-- name: VagueStockSearch :many

SELECT * FROM stock 
WHERE LOWER(stockname) LIKE LOWER($1);