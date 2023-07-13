
-- name: CreateCustomers :one
INSERT INTO customers (
    name, address, contact_number, email
) VALUES (
 $1, $2, $3, $4
)