-- name: CreateSales_till :one
INSERT INTO sales_till (
    till_num, teller, supervisor, branch, close_time, close_cash, close_summary
) VALUES (
     $1, $2, $3, $4, $5, $6, $7
)

-- name: GetSales_till :one
SELECT * FROM sales_till
WHERE till_num = $1 LIMIT 1;