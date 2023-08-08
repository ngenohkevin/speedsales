-- name: CreateSales :one
INSERT INTO sales (
        receipt_num, till_num, product_id, item_name, vat_code, hs_code, batch_code, serial_code, serial_code_return, served_by, approved_by
) VALUES (
        $1,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetSales :one
SELECT * FROM sales
WHERE receipt_num = $1 LIMIT 1;

-- name: ListSales :many
SELECT * FROM sales
WHERE receipt_num = $1
ORDER BY item_name
LIMIT $2
OFFSET $3;

-- name: UpdateSale :one
UPDATE sales
SET item_name = $2,
    vat_code = $3,
    hs_code = $4,
    batch_code = $5,
    serial_code = $6,
    serial_code_return = $7
WHERE product_id = $1
RETURNING *;

-- name: DeleteSale :exec
DELETE FROM sales
WHERE product_id = $1;



