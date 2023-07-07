-- name: CreateProducts :one
INSERT INTO products (
    name, description, category, department_id, supplier_id, cost, selling_price, wholesale_price, min_margin, quantity
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

