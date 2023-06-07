-- name: CreateCustomer :one
INSERT INTO customers (
    name,
    address,
    contact_number,
    email
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;


-- name: GetCustomer :one
SELECT * FROM customers
WHERE customer_id = $1 LIMIT 1;


-- name: ListCustomer :many
SELECT * FROM customers
ORDER BY customer_id
LIMIT $1
    OFFSET $2;

-- name: UpdateCustomer :one
UPDATE customers
SET  name = $2,
     address = $3,
     contact_number = $4,
     email = $5
WHERE customer_id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE customer_id = $1;

