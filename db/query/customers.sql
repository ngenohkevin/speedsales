
-- name: CreateCustomers :one
INSERT INTO customers (
    name, address, contact_number, email
) VALUES (
 $1, $2, $3, $4
) RETURNING *;

-- name: GetCustomers :one
SELECT * FROM customers
WHERE customer_id = $1
ORDER BY name
LIMIT $2
OFFSET $3;