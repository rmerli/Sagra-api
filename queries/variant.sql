-- name: GetVariant :one
SELECT * FROM variants
WHERE id = $1 LIMIT 1;

-- name: ListVariants :many
SELECT * FROM variants
ORDER BY name;

-- name: CreateVariant :one
INSERT INTO variants (
  name, price
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateVariant :one
UPDATE variants
  set name = $2, price = $3
  WHERE id = $1
RETURNING *;

-- name: DeleteVariant :exec
DELETE FROM variants
WHERE id = $1;
