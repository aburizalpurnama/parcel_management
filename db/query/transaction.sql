-- name: CreateTransaction :one
INSERT INTO transactions (
unit_id, delivered_by, type, qty, owner, phone, user_in_id 
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetTransactionById :one
SELECT * FROM transactions
WHERE id = $1 and deleted_at IS NULL;

-- name: GetListAllTransactions :many
SELECT * FROM transactions
WHERE deleted_at IS NULL
ORDER BY id;

-- name: GetListPendingTransactions :many
SELECT * FROM transactions
WHERE deleted_at IS NULL
AND picked_at IS NULL
ORDER BY id;

-- name: GetListDoneTransactions :many
SELECT * FROM transactions
WHERE deleted_at IS NULL
AND picked_at IS NOT NULL
ORDER BY id;

-- name: UpdateTransaction :exec
UPDATE transactions 
SET qty = $1, user_out_id = $2, picked_by = $3, picked_at = (now()), updated_at = (now())
WHERE id = $4;

-- name: DeleteTransaction :exec
UPDATE transactions SET deleted_at = (now())
WHERE id = $1;