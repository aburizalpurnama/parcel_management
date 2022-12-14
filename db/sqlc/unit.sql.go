// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: unit.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUnit = `-- name: CreateUnit :one
INSERT INTO units (
    no, email, phone
) VALUES (
    $1, $2, $3
)
RETURNING id, no, email, item_pending_qty, phone, deleted_at, created_at, updated_at
`

type CreateUnitParams struct {
	No    string `json:"no"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (q *Queries) CreateUnit(ctx context.Context, arg CreateUnitParams) (Unit, error) {
	row := q.db.QueryRowContext(ctx, createUnit, arg.No, arg.Email, arg.Phone)
	var i Unit
	err := row.Scan(
		&i.ID,
		&i.No,
		&i.Email,
		&i.ItemPendingQty,
		&i.Phone,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUnitById = `-- name: DeleteUnitById :exec
UPDATE units SET deleted_at = (now())
WHERE id = $1
`

func (q *Queries) DeleteUnitById(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUnitById, id)
	return err
}

const getListUnits = `-- name: GetListUnits :many
SELECT id, no, email, item_pending_qty, phone, deleted_at, created_at, updated_at FROM units
WHERE deleted_at IS NULL
ORDER BY id
`

func (q *Queries) GetListUnits(ctx context.Context) ([]Unit, error) {
	rows, err := q.db.QueryContext(ctx, getListUnits)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Unit
	for rows.Next() {
		var i Unit
		if err := rows.Scan(
			&i.ID,
			&i.No,
			&i.Email,
			&i.ItemPendingQty,
			&i.Phone,
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

const getUnitById = `-- name: GetUnitById :one
SELECT id, no, email, item_pending_qty, phone, deleted_at, created_at, updated_at FROM units
WHERE id = $1 and deleted_at IS NULL
`

func (q *Queries) GetUnitById(ctx context.Context, id uuid.UUID) (Unit, error) {
	row := q.db.QueryRowContext(ctx, getUnitById, id)
	var i Unit
	err := row.Scan(
		&i.ID,
		&i.No,
		&i.Email,
		&i.ItemPendingQty,
		&i.Phone,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePhoneUnit = `-- name: UpdatePhoneUnit :exec
UPDATE units SET phone = $1, updated_at = (now()) 
WHERE id = $2
`

type UpdatePhoneUnitParams struct {
	Phone string    `json:"phone"`
	ID    uuid.UUID `json:"id"`
}

func (q *Queries) UpdatePhoneUnit(ctx context.Context, arg UpdatePhoneUnitParams) error {
	_, err := q.db.ExecContext(ctx, updatePhoneUnit, arg.Phone, arg.ID)
	return err
}
