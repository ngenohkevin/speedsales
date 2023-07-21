-- name: CreateSales_Till :one
INSERT INTO sales_till (
    teller, supervisor, branch, open_time, open_cash, close_time, close_cash, close_summary
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;
