-- name: CreateDepartment :one
INSERT INTO department (category, sub_category, description)
VALUES (
 $1, $2, $3
) RETURNING *

-- name: GetDepartment :one
SELECT * FROM department
WHERE department_id = $1 LIMIT 1;

-- name: ListDepartment :many