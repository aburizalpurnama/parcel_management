// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: transaction.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
unit_id, delivered_by, type, qty, owner, phone, user_in_id 
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id, unit_id, delivered_by, type, qty, owner, phone, user_in_id, user_out_id, picked_by, picked_at, deleted_at, created_at, updated_at
`

type CreateTransactionParams struct {
	UnitID      uuid.NullUUID `json:"unit_id"`
	DeliveredBy string        `json:"delivered_by"`
	Type        ProductTypes  `json:"type"`
	Qty         int32         `json:"qty"`
	Owner       string        `json:"owner"`
	Phone       string        `json:"phone"`
	UserInID    uuid.NullUUID `json:"user_in_id"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.UnitID,
		arg.DeliveredBy,
		arg.Type,
		arg.Qty,
		arg.Owner,
		arg.Phone,
		arg.UserInID,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UnitID,
		&i.DeliveredBy,
		&i.Type,
		&i.Qty,
		&i.Owner,
		&i.Phone,
		&i.UserInID,
		&i.UserOutID,
		&i.PickedBy,
		&i.PickedAt,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
UPDATE transactions SET deleted_at = (now())
WHERE id = $1
`

func (q *Queries) DeleteTransaction(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTransaction, id)
	return err
}

const getListTransactions = `-- name: GetListTransactions :many
SELECT id, unit_id, delivered_by, type, qty, owner, phone, user_in_id, user_out_id, picked_by, picked_at, deleted_at, created_at, updated_at FROM transactions
WHERE deleted_at IS NULL
ORDER BY id
`

func (q *Queries) GetListTransactions(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getListTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UnitID,
			&i.DeliveredBy,
			&i.Type,
			&i.Qty,
			&i.Owner,
			&i.Phone,
			&i.UserInID,
			&i.UserOutID,
			&i.PickedBy,
			&i.PickedAt,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getTransactionById = `-- name: GetTransactionById :one
SELECT id, unit_id, delivered_by, type, qty, owner, phone, user_in_id, user_out_id, picked_by, picked_at, deleted_at, created_at, updated_at FROM transactions
WHERE id = $1 and deleted_at IS NULL
`

func (q *Queries) GetTransactionById(ctx context.Context, id uuid.UUID) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransactionById, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UnitID,
		&i.DeliveredBy,
		&i.Type,
		&i.Qty,
		&i.Owner,
		&i.Phone,
		&i.UserInID,
		&i.UserOutID,
		&i.PickedBy,
		&i.PickedAt,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTransaction = `-- name: UpdateTransaction :exec
UPDATE transactions 
SET qty = $1, user_out_id = $2, picked_by = $3, picked_at = (now()), updated_at = (now())
WHERE id = $2
`

type UpdateTransactionParams struct {
	Qty       int32         `json:"qty"`
	UserOutID uuid.NullUUID `json:"user_out_id"`
	PickedBy  string        `json:"picked_by"`
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) error {
	_, err := q.db.ExecContext(ctx, updateTransaction, arg.Qty, arg.UserOutID, arg.PickedBy)
	return err
}
