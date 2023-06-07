-- name: CreateSale :one
INSERT INTO sales (
   product_id, customer_id, quantity, sale_date, total_price

) VALUES (
             $1, $2, $3, $4, $5
         ) RETURNING *;


-- name: GetSale :one
SELECT * FROM sales
WHERE sale_id = $1 LIMIT 1;


-- name: ListSale :many
SELECT * FROM sales
ORDER BY sale_id
LIMIT $1
    OFFSET $2;

-- name: UpdateSale :one
UPDATE sales
SET  quantity = $2,
     sale_date = $3,
     total_price = $4
WHERE sale_id = $1
RETURNING *;

-- name: DeleteSale :exec
DELETE FROM sales
WHERE sale_id = $1;

