-- name: GetPassword :one

SELECT  pass_word FROM users
WHERE user_name = $1;