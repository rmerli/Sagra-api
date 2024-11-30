package model

import (
	"gtmx/src/database"

	"github.com/jackc/pgx/v5/pgtype"
)

type Menu struct {
	Id    int64
	Name  string
	Start pgtype.Date
	End   pgtype.Date
}

func NewMenu(id int64, name string, start pgtype.Date, end pgtype.Date) Menu {
	return Menu{
		Id:    id,
		Name:  name,
		Start: start,
		End:   end,
	}
}

func NewMenuList(dbMenus []database.Menu) []Menu {
	menus := make([]Menu, len(dbMenus))

	for i, menu := range dbMenus {
		menus[i] = Menu{
			Id:    menu.ID,
			Name:  menu.Name,
			Start: menu.StartDate,
			End:   menu.EndDate,
		}
	}

	return menus
}
