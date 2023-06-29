
-- name: CreateUser :one
INSERT INTO users (
    username, branch, stk_location, reset, till_num, rights, is_active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
WHERE user_id = $1
ORDER BY username
LIMIT $2
OFFSET $3;

-- name: UpdateUser :one
UPDATE users
SET username = $2,
    branch = $3,
    stk_location = $4,
    reset = $5,
    till_num =$6,
    rights = $7,
    is_active = $8
WHERE user_id = $1
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE user_id = $1;


