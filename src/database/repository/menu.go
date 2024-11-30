package repository

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/model"
)

type MenuRepository struct {
	db *database.Queries
}

func (r MenuRepository) List(ctx context.Context) ([]model.Menu, error) {
	menus, err := r.db.ListMenus(ctx)
	if err != nil {
		return []model.Menu{}, err

	}

	return model.NewMenuList(menus), nil
}

func (r MenuRepository) Get(ctx context.Context, id int64) (model.Menu, error) {
	menu, err := r.db.GetMenu(ctx, id)

	if err != nil {
		return model.Menu{}, err

	}

	return model.NewMenu(menu.ID, menu.Name, menu.StartDate, menu.EndDate), nil
}

func (r MenuRepository) Insert(ctx context.Context, menu model.Menu) (model.Menu, error) {

	insertedMenu, err := r.db.CreateMenu(ctx, database.CreateMenuParams{
		Name:      menu.Name,
		StartDate: menu.Start,
		EndDate:   menu.End,
	})
	if err != nil {
		return model.Menu{}, err
	}

	menu.Id = insertedMenu.ID

	return menu, nil
}

func (r MenuRepository) Update(ctx context.Context, menu model.Menu) (model.Menu, error) {

	_, err := r.db.UpdateMenu(ctx, database.UpdateMenuParams{
		ID:        menu.Id,
		Name:      menu.Name,
		StartDate: menu.Start,
		EndDate:   menu.End,
	})
	if err != nil {
		return model.Menu{}, err
	}

	return menu, nil
}

func NewMenuRepository(db *database.Queries) MenuRepository {
	return MenuRepository{db: db}
}
