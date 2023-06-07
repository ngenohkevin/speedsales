-- name: CreateProduct :one
INSERT INTO products (
    name,
    description,
    category,
    supplier_id,
    cost,
    selling_price,
    quantity
) VALUES (
             $1, $2, $3, $4, $5, $6, $7
         ) RETURNING *;


-- name: GetProduct :one
SELECT * FROM products
WHERE product_id = $1 LIMIT 1;


-- name: ListProduct :many
SELECT * FROM products
ORDER BY product_id
LIMIT $1
    OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
SET  name = $2,
     description = $3,
     category = $4,
     cost = $5,
     selling_price = $6,
     quantity = $7
WHERE product_id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1;

