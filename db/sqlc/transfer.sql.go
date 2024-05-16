// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO tranfers (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Tranfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Tranfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTranfer = `-- name: GetTranfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM tranfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTranfer(ctx context.Context, id int64) (Tranfer, error) {
	row := q.db.QueryRowContext(ctx, getTranfer, id)
	var i Tranfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTranfers = `-- name: ListTranfers :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM tranfers
WHERE 
  from_account_id = $1 OR
  to_account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTranfersParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTranfers(ctx context.Context, arg ListTranfersParams) ([]Tranfer, error) {
	rows, err := q.db.QueryContext(ctx, listTranfers,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tranfer
	for rows.Next() {
		var i Tranfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
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
