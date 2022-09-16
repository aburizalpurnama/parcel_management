-- name: CreateUser :one
INSERT INTO users (
    name, email, password, role_id
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 and deleted_at IS NULL;

-- name: GetListUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL
ORDER BY id;

-- name: UpdatePasswordUser :exec
UPDATE users SET password = $1 
WHERE id = $2;

-- name: DeleteUserById :exec
UPDATE users SET deleted_at = (now())
WHERE id = $1;