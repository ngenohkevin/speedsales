-- name: CreateSupplier :one
INSERT INTO suppliers (
   name, address, contact_number, email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetSupplier :one
SELECT * FROM suppliers
WHERE supplier_id = $1 LIMIT 1;

-- name: ListSuppliers :many
SELECT * FROM suppliers
WHERE supplier_id = $1
ORDER BY name
LIMIT $2
OFFSET $3;

-- name: UpdateSupplier :one
UPDATE suppliers
SET name = $2,
    address = $3,
    contact_number = $4,
    email = $5
WHERE supplier_id = $1
RETURNING *;

-- name: DeleteSupplier :exec
DELETE FROM suppliers
WHERE supplier_id = $1;