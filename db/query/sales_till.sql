-- name: CreateSales_till :one
INSERT INTO sales_till (
    till_num, teller, supervisor, branch, open_cash, close_time, close_cash, close_summary
) VALUES (
     $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetSales_till :one
SELECT * FROM sales_till
WHERE till_num = $1 LIMIT 1;

-- name: ListSales_till :many
SELECT * FROM sales_till
WHERE till_num = $1
ORDER BY teller
LIMIT $2
OFFSET $3;

-- name: UpdateSales_till :one
UPDATE sales_till
SET teller = $2,
    supervisor = $3,
    branch = $4,
    open_cash = $5,
    close_time = $6,
    close_time = $7,
    close_cash = $8,
    close_summary = $9
WHERE till_num = $1
RETURNING *;

-- name: DeleteSales_till :exec
DELETE FROM sales_till
WHERE till_num = $1;