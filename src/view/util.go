package view

import (
	"fmt"
	"gtmx/src/server/routes"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func PathReplaceId(path string, id uuid.UUID) string {
	return strings.Replace(routes.GetPath(path), ":id", id.String(), 1)
}

func FormatPrice(price pgtype.Numeric) string {
	p, err := price.Float64Value()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%.2f", p.Float64)
}
