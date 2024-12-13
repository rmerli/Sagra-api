package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Menu struct {
	Model
	Name       string         `json:"name"`
	StartDate  pgtype.Date    `json:"startDate"`
	EndDate    pgtype.Date    `json:"endDate"`
	Categories []MenuCategory `json:"categories"`
}
