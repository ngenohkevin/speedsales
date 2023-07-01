-- name: CreateCodeTranslator :one
INSERT INTO code_translator (
  master_code, link_code, pkg_qty, discount
) VALUES (
  $1, $2, $3, $4
)