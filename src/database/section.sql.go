// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: section.sql

package database

import (
	"context"
)

const createSection = `-- name: CreateSection :one
INSERT INTO sections (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateSection(ctx context.Context, name string) (Section, error) {
	row := q.db.QueryRow(ctx, createSection, name)
	var i Section
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteSection = `-- name: DeleteSection :exec
DELETE FROM sections
WHERE id = $1
`

func (q *Queries) DeleteSection(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteSection, id)
	return err
}

const getSection = `-- name: GetSection :one
SELECT id, name FROM sections
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSection(ctx context.Context, id int64) (Section, error) {
	row := q.db.QueryRow(ctx, getSection, id)
	var i Section
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listSections = `-- name: ListSections :many
SELECT id, name FROM sections
ORDER BY name
`

func (q *Queries) ListSections(ctx context.Context) ([]Section, error) {
	rows, err := q.db.Query(ctx, listSections)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Section
	for rows.Next() {
		var i Section
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSection = `-- name: UpdateSection :one
UPDATE sections
  set name = $2
  WHERE id = $1
RETURNING id, name
`

type UpdateSectionParams struct {
	ID   int64
	Name string
}

func (q *Queries) UpdateSection(ctx context.Context, arg UpdateSectionParams) (Section, error) {
	row := q.db.QueryRow(ctx, updateSection, arg.ID, arg.Name)
	var i Section
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}