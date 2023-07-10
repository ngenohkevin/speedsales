// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: suppliers.sql

package db

import (
	"context"
)

const createSupplier = `-- name: CreateSupplier :one
INSERT INTO suppliers (
   name, address, contact_number, email
) VALUES (
    $1, $2, $3, $4
) RETURNING supplier_id, name, address, contact_number, email, created_at
`

type CreateSupplierParams struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	ContactNumber string `json:"contact_number"`
	Email         string `json:"email"`
}

func (q *Queries) CreateSupplier(ctx context.Context, arg CreateSupplierParams) (Supplier, error) {
	row := q.db.QueryRow(ctx, createSupplier,
		arg.Name,
		arg.Address,
		arg.ContactNumber,
		arg.Email,
	)
	var i Supplier
	err := row.Scan(
		&i.SupplierID,
		&i.Name,
		&i.Address,
		&i.ContactNumber,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSupplier = `-- name: DeleteSupplier :exec
DELETE FROM suppliers
WHERE supplier_id = $1
`

func (q *Queries) DeleteSupplier(ctx context.Context, supplierID int32) error {
	_, err := q.db.Exec(ctx, deleteSupplier, supplierID)
	return err
}

const getSupplier = `-- name: GetSupplier :one
SELECT supplier_id, name, address, contact_number, email, created_at FROM suppliers
WHERE supplier_id = $1 LIMIT 1
`

func (q *Queries) GetSupplier(ctx context.Context, supplierID int32) (Supplier, error) {
	row := q.db.QueryRow(ctx, getSupplier, supplierID)
	var i Supplier
	err := row.Scan(
		&i.SupplierID,
		&i.Name,
		&i.Address,
		&i.ContactNumber,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const listSuppliers = `-- name: ListSuppliers :many
SELECT supplier_id, name, address, contact_number, email, created_at FROM suppliers
WHERE supplier_id = $1
ORDER BY name
LIMIT $2
OFFSET $3
`

type ListSuppliersParams struct {
	SupplierID int32 `json:"supplier_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListSuppliers(ctx context.Context, arg ListSuppliersParams) ([]Supplier, error) {
	rows, err := q.db.Query(ctx, listSuppliers, arg.SupplierID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Supplier{}
	for rows.Next() {
		var i Supplier
		if err := rows.Scan(
			&i.SupplierID,
			&i.Name,
			&i.Address,
			&i.ContactNumber,
			&i.Email,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSupplier = `-- name: UpdateSupplier :one
UPDATE suppliers
SET name = $2,
    address = $3,
    contact_number = $4,
    email = $5
WHERE supplier_id = $1
RETURNING supplier_id, name, address, contact_number, email, created_at
`

type UpdateSupplierParams struct {
	SupplierID    int32  `json:"supplier_id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	ContactNumber string `json:"contact_number"`
	Email         string `json:"email"`
}

func (q *Queries) UpdateSupplier(ctx context.Context, arg UpdateSupplierParams) (Supplier, error) {
	row := q.db.QueryRow(ctx, updateSupplier,
		arg.SupplierID,
		arg.Name,
		arg.Address,
		arg.ContactNumber,
		arg.Email,
	)
	var i Supplier
	err := row.Scan(
		&i.SupplierID,
		&i.Name,
		&i.Address,
		&i.ContactNumber,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}
