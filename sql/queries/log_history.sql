-- name: LogHistory :exec

INSERT INTO history(id, user_name, stockname, edit_time, adjustment)
VALUES(gen_random_uuid(), $1, $2, NOW(), $3);