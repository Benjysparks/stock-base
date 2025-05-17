-- name: ShowStockHistory :many

SELECT  TO_CHAR(edit_time, 'DD-MM-YYYY HH24:MI:SS'), user_name, adjustment FROM history
WHERE stockname = $1
ORDER BY edit_time DESC;