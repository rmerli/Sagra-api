-- name: GetMenu :one
SELECT * FROM menus
WHERE id = $1 LIMIT 1;

-- name: ListMenus :many
SELECT * FROM menus
ORDER BY start_date;

-- name: CreateMenu :one
INSERT INTO menus (
  name, start_date, end_date
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateMenu :one
UPDATE menus
set name = $2, start_date = $3, end_date = $4
  WHERE id = $1
RETURNING *;

-- name: DeleteMenu :exec
DELETE FROM menus
WHERE id = $1;
