package view

import (
	"fmt"
	"gtmx/src/server/routes"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

func PathReplaceId(path string, id int64) string {
	return strings.Replace(routes.GetPath(path), ":id", fmt.Sprint(id), 1)
}

func FormatPrice(price pgtype.Numeric) string {
	p, err := price.Float64Value()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%.2f", p.Float64)
}
