-- name: CreateCodeTranslator :one
INSERT INTO code_translator (
  master_code, link_code, pkg_qty, discount
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetCodeTranslator :one
SELECT * FROM code_translator
WHERE master_code = $1 LIMIT 1;

-- name: ListCodeTranslator :many
SELECT * FROM code_translator
WHERE master_code = $1
ORDER BY link_code
LIMIT $2
OFFSET $3;
