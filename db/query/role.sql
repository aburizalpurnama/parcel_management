-- name: CreateRole :one
INSERT INTO roles (
    name
) VALUES (
    $1
)
RETURNING *;

-- name: GetRoleById :one
SELECT * FROM roles
WHERE id = $1 and deleted_at IS NULL;

-- name: GetListRoles :many
SELECT * FROM roles
WHERE deleted_at IS NULL
ORDER BY id;

-- name: UpdateNameRole :exec
UPDATE roles SET name = $1 
WHERE id = $2;

-- name: DeleteRoleById :exec
UPDATE roles SET deleted_at = (now())
WHERE id = $1;