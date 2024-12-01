-- name: GetSection :one
SELECT * FROM sections
WHERE id = $1;

-- name: ListSections :many
SELECT * FROM sections
ORDER BY name;

-- name: GetSectionWithCategories :many
SELECT sqlc.embed(sections), sqlc.embed(categories)
FROM sections
INNER JOIN categories ON categories.section_id = sections.id
WHERE sections.id = $1;

-- name: CreateSection :one
INSERT INTO sections (name)
VALUES ($1)
RETURNING *;

-- name: UpdateSection :one
UPDATE sections
  set name = $2
  WHERE id = $1
RETURNING *;

-- name: DeleteSection :exec
DELETE FROM sections
WHERE id = $1;
