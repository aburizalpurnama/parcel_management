-- name: CreateUnit :one
INSERT INTO units (
    no, email, phone
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetUnitById :one
SELECT * FROM units
WHERE id = $1 and deleted_at IS NULL;

-- name: GetListUnits :many
SELECT * FROM units
WHERE deleted_at IS NULL
ORDER BY id;

-- name: UpdatePhoneUnit :exec
UPDATE units SET phone = $1, updated_at = (now()) 
WHERE id = $2;

-- name: DeleteUnitById :exec
UPDATE units SET deleted_at = (now())
WHERE id = $1;