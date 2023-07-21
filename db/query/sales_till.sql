-- name: CreateSales_Till :one
INSERT INTO sales_till (
    till_num, teller, supervisor, branch, close_time, close_cash, close_summary
) VALUES (
     $1, $2, $3, $4, $5, $6, $7
)
