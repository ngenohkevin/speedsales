// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: customers.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customers (
    name, address, contact_number, email
) VALUES (
 $1, $2, $3, $4
) RETURNING customer_id, name, address, contact_number, email, created_at
`

type CreateCustomerParams struct {
	Name          string      `json:"name"`
	Address       pgtype.Text `json:"address"`
	ContactNumber string      `json:"contact_number"`
	Email         pgtype.Text `json:"email"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRow(ctx, createCustomer,
		arg.Name,
		arg.Address,
		arg.ContactNumber,
		arg.Email,
	)
	var i Customer
	err := row.Scan(
		&i.CustomerID,
		&i.Name,
		&i.Address,
		&i.ContactNumber,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE customer_id = $1
`

func (q *Queries) DeleteCustomer(ctx context.Context, customerID int64) error {
	_, err := q.db.Exec(ctx, deleteCustomer, customerID)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT customer_id, name, address, contact_number, email, created_at FROM customers
WHERE customer_id = $1 LIMIT 1
`

func (q *Queries) GetCustomer(ctx context.Context, customerID int64) (Customer, error) {
	row := q.db.QueryRow(ctx, getCustomer, customerID)
	var i Customer
	err := row.Scan(
		&i.CustomerID,
		&i.Name,
		&i.Address,
		&i.ContactNumber,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const listCustomers = `-- name: ListCustomers :many
SELECT customer_id, name, address, contact_number, email, created_at FROM customers
WHERE customer_id = $1
ORDER BY name
LIMIT $2
OFFSET $3
`

type ListCustomersParams struct {
	CustomerID int64 `json:"customer_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListCustomers(ctx context.Context, arg ListCustomersParams) ([]Customer, error) {
	rows, err := q.db.Query(ctx, listCustomers, arg.CustomerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Customer{}
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.CustomerID,
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

const updateCustomer = `-- name: UpdateCustomer :one
UPDATE customers
SET name = $2,
    address = $3,
    contact_number = $4,
    email = $5
WHERE customer_id = $1
RETURNING customer_id, name, address, contact_number, email, created_at
`

type UpdateCustomerParams struct {
	CustomerID    int64       `json:"customer_id"`
	Name          string      `json:"name"`
	Address       pgtype.Text `json:"address"`
	ContactNumber string      `json:"contact_number"`
	Email         pgtype.Text `json:"email"`
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error) {
	row := q.db.QueryRow(ctx, updateCustomer,
		arg.CustomerID,
		arg.Name,
		arg.Address,
		arg.ContactNumber,
		arg.Email,
	)
	var i Customer
	err := row.Scan(
		&i.CustomerID,
		&i.Name,
		&i.Address,
		&i.ContactNumber,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}