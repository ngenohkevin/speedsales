-- name: CreateProducts :one
INSERT INTO products (
    name, description, category, department_id, supplier_id, cost, selling_price, wholesale_price, min_margin, quantity
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetProducts :one
SELECT * FROM products
WHERE product_id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
WHERE product_id = $1
ORDER BY name
LIMIT $2
OFFSET $3;