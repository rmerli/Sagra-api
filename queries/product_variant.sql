-- name: CreateProductVariant :one
INSERT INTO products_variants (product_id, variant_id)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteProductVariant :exec
DELETE FROM products_variants
WHERE product_id = $1 AND variant_id = $2;

-- name: RemoveAllVariantToProduct :exec
DELETE FROM products_variants
WHERE product_id = $1;

