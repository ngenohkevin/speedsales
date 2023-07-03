-- name: CreateDepartment :one
INSERT INTO department (category, sub_category, description)
VALUES (
 $1, $2, $3
) RETURNING *

-- name: GetDepartment :one
SELECT * FROM department
WHERE department_id = $1 LIMIT 1;

-- name: ListDepartment :many
SELECT * FROM department
WHERE department_id = $1
ORDER BY category
LIMIT $2
OFFSET $3;

-- name UpdateDepartment :one
UPDATE department
SET category = $2,
    sub_category = $3,
    description = $4
WHERE department_id = $1
RETURNING *;

-- name: DeleteDepartment :exec
DELETE FROM department
WHERE department_id = $1;