-- name: CreateSales_till :one
INSERT INTO sales_till (
    till_num, teller, supervisor, branch, close_time, close_cash, close_summary
) VALUES (
     $1, $2, $3, $4, $5, $6, $7
);

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
SET till_num = $2,
    teller = $3,
    supervisor = $4,
    open_cash = $5,
    close_time = $6,
    close_cash = $7,
    close_time = $8
WHERE till_num = $1
RETURNING *;

