-- name: GetInvoiceLines :many

SELECT  stockname, amount, qty_type, price_per, total_price, TO_CHAR(created_at, 'DD-MM-YYYY'), user_name FROM invoice;