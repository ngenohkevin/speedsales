// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: products.sql

package db

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
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
         ) RETURNING product_id, name, description, category, supplier_id, cost, selling_price, quantity
`

type CreateProductParams struct {
	Name         sql.NullString `json:"name"`
	Description  sql.NullString `json:"description"`
	Category     sql.NullString `json:"category"`
	SupplierID   sql.NullInt32  `json:"supplier_id"`
	Cost         sql.NullString `json:"cost"`
	SellingPrice sql.NullString `json:"selling_price"`
	Quantity     sql.NullInt32  `json:"quantity"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Name,
		arg.Description,
		arg.Category,
		arg.SupplierID,
		arg.Cost,
		arg.SellingPrice,
		arg.Quantity,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Description,
		&i.Category,
		&i.SupplierID,
		&i.Cost,
		&i.SellingPrice,
		&i.Quantity,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, productID int32) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, productID)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT product_id, name, description, category, supplier_id, cost, selling_price, quantity FROM products
WHERE product_id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, productID int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Description,
		&i.Category,
		&i.SupplierID,
		&i.Cost,
		&i.SellingPrice,
		&i.Quantity,
	)
	return i, err
}

const listProduct = `-- name: ListProduct :many
SELECT product_id, name, description, category, supplier_id, cost, selling_price, quantity FROM products
ORDER BY product_id
LIMIT $1
    OFFSET $2
`

type ListProductParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProduct(ctx context.Context, arg ListProductParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProduct, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.Name,
			&i.Description,
			&i.Category,
			&i.SupplierID,
			&i.Cost,
			&i.SellingPrice,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET  name = $2,
     description = $3,
     category = $4,
     cost = $5,
     selling_price = $6,
     quantity = $7
WHERE product_id = $1
RETURNING product_id, name, description, category, supplier_id, cost, selling_price, quantity
`

type UpdateProductParams struct {
	ProductID    int32          `json:"product_id"`
	Name         sql.NullString `json:"name"`
	Description  sql.NullString `json:"description"`
	Category     sql.NullString `json:"category"`
	Cost         sql.NullString `json:"cost"`
	SellingPrice sql.NullString `json:"selling_price"`
	Quantity     sql.NullInt32  `json:"quantity"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.ProductID,
		arg.Name,
		arg.Description,
		arg.Category,
		arg.Cost,
		arg.SellingPrice,
		arg.Quantity,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Description,
		&i.Category,
		&i.SupplierID,
		&i.Cost,
		&i.SellingPrice,
		&i.Quantity,
	)
	return i, err
}