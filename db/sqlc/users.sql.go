// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username, branch, stk_location, reset, till_num, rights, is_active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING user_id, username, branch, stk_location, reset, till_num, rights, is_active
`

type CreateUserParams struct {
	Username    pgtype.Text `json:"username"`
	Branch      pgtype.Text `json:"branch"`
	StkLocation pgtype.Text `json:"stk_location"`
	Reset       pgtype.Text `json:"reset"`
	TillNum     pgtype.Int8 `json:"till_num"`
	Rights      []byte      `json:"rights"`
	IsActive    pgtype.Bool `json:"is_active"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Branch,
		arg.StkLocation,
		arg.Reset,
		arg.TillNum,
		arg.Rights,
		arg.IsActive,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Branch,
		&i.StkLocation,
		&i.Reset,
		&i.TillNum,
		&i.Rights,
		&i.IsActive,
	)
	return i, err
}

const deleteUsers = `-- name: DeleteUsers :exec
DELETE FROM users
WHERE user_id = $1
`

func (q *Queries) DeleteUsers(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteUsers, userID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_id, username, branch, stk_location, reset, till_num, rights, is_active FROM users
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userID int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Branch,
		&i.StkLocation,
		&i.Reset,
		&i.TillNum,
		&i.Rights,
		&i.IsActive,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, username, branch, stk_location, reset, till_num, rights, is_active FROM users
WHERE user_id = $1
ORDER BY username
LIMIT $2
OFFSET $3
`

type ListUsersParams struct {
	UserID int32 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Branch,
			&i.StkLocation,
			&i.Reset,
			&i.TillNum,
			&i.Rights,
			&i.IsActive,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET username = $2,
    branch = $3,
    stk_location = $4,
    reset = $5,
    till_num =$6,
    rights = $7,
    is_active = $8
WHERE user_id = $1
RETURNING user_id, username, branch, stk_location, reset, till_num, rights, is_active
`

type UpdateUserParams struct {
	UserID      int32       `json:"user_id"`
	Username    pgtype.Text `json:"username"`
	Branch      pgtype.Text `json:"branch"`
	StkLocation pgtype.Text `json:"stk_location"`
	Reset       pgtype.Text `json:"reset"`
	TillNum     pgtype.Int8 `json:"till_num"`
	Rights      []byte      `json:"rights"`
	IsActive    pgtype.Bool `json:"is_active"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.UserID,
		arg.Username,
		arg.Branch,
		arg.StkLocation,
		arg.Reset,
		arg.TillNum,
		arg.Rights,
		arg.IsActive,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Branch,
		&i.StkLocation,
		&i.Reset,
		&i.TillNum,
		&i.Rights,
		&i.IsActive,
	)
	return i, err
}
