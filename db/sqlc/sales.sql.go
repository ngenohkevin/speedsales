// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: sales.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSales = `-- name: CreateSales :one
INSERT INTO sales (
        receipt_num, till_num, product_id, item_name, vat_code, hs_code, batch_code, serial_code, serial_code_return, served_by, approved_by
) VALUES (
        $1,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING receipt_num, till_num, txn_time, product_id, item_name, price, cost, quantity, vat_code, hs_code, "VAT", batch_code, serial_code, serial_code_return, served_by, approved_by, state
`

type CreateSalesParams struct {
	ReceiptNum       int64       `json:"receipt_num"`
	TillNum          pgtype.Int8 `json:"till_num"`
	ProductID        int64       `json:"product_id"`
	ItemName         pgtype.Text `json:"item_name"`
	VatCode          pgtype.Text `json:"vat_code"`
	HsCode           pgtype.Text `json:"hs_code"`
	BatchCode        pgtype.Text `json:"batch_code"`
	SerialCode       pgtype.Text `json:"serial_code"`
	SerialCodeReturn pgtype.Text `json:"serial_code_return"`
	ServedBy         pgtype.Text `json:"served_by"`
	ApprovedBy       pgtype.Text `json:"approved_by"`
}

func (q *Queries) CreateSales(ctx context.Context, arg CreateSalesParams) (Sale, error) {
	row := q.db.QueryRow(ctx, createSales,
		arg.ReceiptNum,
		arg.TillNum,
		arg.ProductID,
		arg.ItemName,
		arg.VatCode,
		arg.HsCode,
		arg.BatchCode,
		arg.SerialCode,
		arg.SerialCodeReturn,
		arg.ServedBy,
		arg.ApprovedBy,
	)
	var i Sale
	err := row.Scan(
		&i.ReceiptNum,
		&i.TillNum,
		&i.TxnTime,
		&i.ProductID,
		&i.ItemName,
		&i.Price,
		&i.Cost,
		&i.Quantity,
		&i.VatCode,
		&i.HsCode,
		&i.VAT,
		&i.BatchCode,
		&i.SerialCode,
		&i.SerialCodeReturn,
		&i.ServedBy,
		&i.ApprovedBy,
		&i.State,
	)
	return i, err
}

const deleteSale = `-- name: DeleteSale :exec
DELETE FROM sales
WHERE receipt_num = $1
`

func (q *Queries) DeleteSale(ctx context.Context, receiptNum int64) error {
	_, err := q.db.Exec(ctx, deleteSale, receiptNum)
	return err
}

const getSales = `-- name: GetSales :one
SELECT receipt_num, till_num, txn_time, product_id, item_name, price, cost, quantity, vat_code, hs_code, "VAT", batch_code, serial_code, serial_code_return, served_by, approved_by, state FROM sales
WHERE receipt_num = $1 LIMIT 1
`

func (q *Queries) GetSales(ctx context.Context, receiptNum int64) (Sale, error) {
	row := q.db.QueryRow(ctx, getSales, receiptNum)
	var i Sale
	err := row.Scan(
		&i.ReceiptNum,
		&i.TillNum,
		&i.TxnTime,
		&i.ProductID,
		&i.ItemName,
		&i.Price,
		&i.Cost,
		&i.Quantity,
		&i.VatCode,
		&i.HsCode,
		&i.VAT,
		&i.BatchCode,
		&i.SerialCode,
		&i.SerialCodeReturn,
		&i.ServedBy,
		&i.ApprovedBy,
		&i.State,
	)
	return i, err
}

const listSales = `-- name: ListSales :many
SELECT receipt_num, till_num, txn_time, product_id, item_name, price, cost, quantity, vat_code, hs_code, "VAT", batch_code, serial_code, serial_code_return, served_by, approved_by, state FROM sales
WHERE receipt_num = $1
ORDER BY item_name
LIMIT $2
OFFSET $3
`

type ListSalesParams struct {
	ReceiptNum int64 `json:"receipt_num"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListSales(ctx context.Context, arg ListSalesParams) ([]Sale, error) {
	rows, err := q.db.Query(ctx, listSales, arg.ReceiptNum, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Sale{}
	for rows.Next() {
		var i Sale
		if err := rows.Scan(
			&i.ReceiptNum,
			&i.TillNum,
			&i.TxnTime,
			&i.ProductID,
			&i.ItemName,
			&i.Price,
			&i.Cost,
			&i.Quantity,
			&i.VatCode,
			&i.HsCode,
			&i.VAT,
			&i.BatchCode,
			&i.SerialCode,
			&i.SerialCodeReturn,
			&i.ServedBy,
			&i.ApprovedBy,
			&i.State,
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

const updateSale = `-- name: UpdateSale :one
UPDATE sales
SET item_name = $2,
    vat_code = $3,
    hs_code = $4,
    batch_code = $5,
    serial_code = $6,
    serial_code_return = $7
WHERE receipt_num = $1
RETURNING receipt_num, till_num, txn_time, product_id, item_name, price, cost, quantity, vat_code, hs_code, "VAT", batch_code, serial_code, serial_code_return, served_by, approved_by, state
`

type UpdateSaleParams struct {
	ReceiptNum       int64       `json:"receipt_num"`
	ItemName         pgtype.Text `json:"item_name"`
	VatCode          pgtype.Text `json:"vat_code"`
	HsCode           pgtype.Text `json:"hs_code"`
	BatchCode        pgtype.Text `json:"batch_code"`
	SerialCode       pgtype.Text `json:"serial_code"`
	SerialCodeReturn pgtype.Text `json:"serial_code_return"`
}

func (q *Queries) UpdateSale(ctx context.Context, arg UpdateSaleParams) (Sale, error) {
	row := q.db.QueryRow(ctx, updateSale,
		arg.ReceiptNum,
		arg.ItemName,
		arg.VatCode,
		arg.HsCode,
		arg.BatchCode,
		arg.SerialCode,
		arg.SerialCodeReturn,
	)
	var i Sale
	err := row.Scan(
		&i.ReceiptNum,
		&i.TillNum,
		&i.TxnTime,
		&i.ProductID,
		&i.ItemName,
		&i.Price,
		&i.Cost,
		&i.Quantity,
		&i.VatCode,
		&i.HsCode,
		&i.VAT,
		&i.BatchCode,
		&i.SerialCode,
		&i.SerialCodeReturn,
		&i.ServedBy,
		&i.ApprovedBy,
		&i.State,
	)
	return i, err
}
