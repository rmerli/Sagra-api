-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: CreateProduct :one
INSERT INTO products (
  name, 
  abbr,
  price
) VALUES (
  $1,$2,$3
)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
  set name = $2,
  price = $3,
  abbr = $4
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

