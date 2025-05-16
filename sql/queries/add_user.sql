-- name: CreateNewUser :one
INSERT INTO users (id, user_name, pass_word)
VALUES (
    gen_random_uuid(),
    $1,
    $2
)
RETURNING *;