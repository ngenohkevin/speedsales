// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: code_translator.sql

package db

import (
	"context"
)

const createCodeTranslator = `-- name: CreateCodeTranslator :one
INSERT INTO code_translator (
  master_code, link_code, pkg_qty, discount
) VALUES (
  $1, $2, $3, $4
) RETURNING master_code, link_code, pkg_qty, discount
`

type CreateCodeTranslatorParams struct {
	MasterCode string  `json:"master_code"`
	LinkCode   string  `json:"link_code"`
	PkgQty     float64 `json:"pkg_qty"`
	Discount   float64 `json:"discount"`
}

func (q *Queries) CreateCodeTranslator(ctx context.Context, arg CreateCodeTranslatorParams) (CodeTranslator, error) {
	row := q.db.QueryRow(ctx, createCodeTranslator,
		arg.MasterCode,
		arg.LinkCode,
		arg.PkgQty,
		arg.Discount,
	)
	var i CodeTranslator
	err := row.Scan(
		&i.MasterCode,
		&i.LinkCode,
		&i.PkgQty,
		&i.Discount,
	)
	return i, err
}

const deleteCodeTranslator = `-- name: DeleteCodeTranslator :exec
DELETE FROM code_translator
WHERE master_code = $1
`

func (q *Queries) DeleteCodeTranslator(ctx context.Context, masterCode string) error {
	_, err := q.db.Exec(ctx, deleteCodeTranslator, masterCode)
	return err
}

const getCodeTranslator = `-- name: GetCodeTranslator :one
SELECT master_code, link_code, pkg_qty, discount FROM code_translator
WHERE master_code = $1 LIMIT 1
`

func (q *Queries) GetCodeTranslator(ctx context.Context, masterCode string) (CodeTranslator, error) {
	row := q.db.QueryRow(ctx, getCodeTranslator, masterCode)
	var i CodeTranslator
	err := row.Scan(
		&i.MasterCode,
		&i.LinkCode,
		&i.PkgQty,
		&i.Discount,
	)
	return i, err
}

const listCodeTranslator = `-- name: ListCodeTranslator :many
SELECT master_code, link_code, pkg_qty, discount FROM code_translator
WHERE master_code = $1
ORDER BY link_code
LIMIT $2
OFFSET $3
`

type ListCodeTranslatorParams struct {
	MasterCode string `json:"master_code"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

func (q *Queries) ListCodeTranslator(ctx context.Context, arg ListCodeTranslatorParams) ([]CodeTranslator, error) {
	rows, err := q.db.Query(ctx, listCodeTranslator, arg.MasterCode, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CodeTranslator{}
	for rows.Next() {
		var i CodeTranslator
		if err := rows.Scan(
			&i.MasterCode,
			&i.LinkCode,
			&i.PkgQty,
			&i.Discount,
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

const updateCodeTranslator = `-- name: UpdateCodeTranslator :one
UPDATE code_translator
SET  pkg_qty = $2,
     discount = $3
WHERE master_code = $1
RETURNING master_code, link_code, pkg_qty, discount
`

type UpdateCodeTranslatorParams struct {
	MasterCode string  `json:"master_code"`
	PkgQty     float64 `json:"pkg_qty"`
	Discount   float64 `json:"discount"`
}

func (q *Queries) UpdateCodeTranslator(ctx context.Context, arg UpdateCodeTranslatorParams) (CodeTranslator, error) {
	row := q.db.QueryRow(ctx, updateCodeTranslator, arg.MasterCode, arg.PkgQty, arg.Discount)
	var i CodeTranslator
	err := row.Scan(
		&i.MasterCode,
		&i.LinkCode,
		&i.PkgQty,
		&i.Discount,
	)
	return i, err
}
