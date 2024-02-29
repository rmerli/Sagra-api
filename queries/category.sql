-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY name;

-- name: CreateCategory :one
INSERT INTO categories (name, section_id)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateCategory :one
UPDATE categories
  set name = $2, section_id = $3
  WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;
