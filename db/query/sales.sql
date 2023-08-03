-- name: CreateSales :one
INSERT INTO sales (
        receipt_num, till_num, product_id, item_name, vat_code, hs_code, batch_code, serial_code, serial_code_return, served_by, approved_by
) VALUES (
        $1,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetSales :one
SELECT * FROM sales
WHERE product_id = $1 LIMIT 1

-- name: ListSale :many