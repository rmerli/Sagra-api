package repository

import (
	"context"
	"gtmx/src/database"
)

type CatalogRepository struct {
	Db *database.Queries
}

func (r CatalogRepository) ListCategories(ctx context.Context) ([]database.Category, error) {
	categories, err := r.Db.ListCategories(ctx)
	if err != nil {
		return []database.Category{}, err

	}

	return categories, nil
}

func (r CatalogRepository) GetOneCategoryById(ctx context.Context, id int64) (database.Category, error) {
	product, err := r.Db.GetCategory(ctx, id)

	if err != nil {
		return database.Category{}, err

	}

	return product, nil
}

func (r CatalogRepository) InsertCategory(ctx context.Context, category database.Category) (database.Category, error) {

	insertedCategory, err := r.Db.CreateCategory(ctx, database.CreateCategoryParams{
		Name:      category.Name,
		SectionID: category.SectionID,
	})
	if err != nil {
		return database.Category{}, err
	}

	return insertedCategory, nil
}

func (r CatalogRepository) ListSections(ctx context.Context) ([]database.Section, error) {
	sections, err := r.Db.ListSections(ctx)
	if err != nil {
		return []database.Section{}, err

	}

	return sections, nil
}

func (r CatalogRepository) GetOneSectionById(ctx context.Context, id int64) (database.Section, error) {
	section, err := r.Db.GetSection(ctx, id)

	if err != nil {
		return database.Section{}, err

	}

	return section, nil
}

func (r CatalogRepository) InsertSection(ctx context.Context, section database.Section) (database.Section, error) {

	insertedSection, err := r.Db.CreateSection(ctx, section.Name)
	if err != nil {
		return database.Section{}, err
	}

	return insertedSection, nil
}

func (r CatalogRepository) ListProducts(ctx context.Context) ([]database.Product, error) {
	products, err := r.Db.ListProducts(ctx)
	if err != nil {
		return []database.Product{}, err

	}

	return products, nil
}

func (r CatalogRepository) GetOneProductById(ctx context.Context, id int64) (database.Product, error) {
	product, err := r.Db.GetProduct(ctx, id)

	if err != nil {
		return database.Product{}, err

	}

	return product, nil
}

func (r CatalogRepository) InsertProduct(ctx context.Context, product database.Product) (database.Product, error) {

	insertedProduct, err := r.Db.CreateProduct(ctx, database.CreateProductParams{
		Name:  product.Name,
		Abbr:  product.Abbr,
		Price: product.Price,
	})
	if err != nil {
		return database.Product{}, err
	}

	return insertedProduct, nil
}

func (r CatalogRepository) ListVariants(ctx context.Context) ([]database.Variant, error) {
	variants, err := r.Db.ListVariants(ctx)
	if err != nil {
		return []database.Variant{}, err

	}

	return variants, nil
}

func (r CatalogRepository) GetOneVariantById(ctx context.Context, id int64) (database.Variant, error) {
	variant, err := r.Db.GetVariant(ctx, id)

	if err != nil {
		return database.Variant{}, err

	}

	return variant, nil
}

func (r CatalogRepository) InsertVariant(ctx context.Context, variant database.Variant) (database.Variant, error) {

	insertedVariant, err := r.Db.CreateVariant(ctx, database.CreateVariantParams{
		Name:  variant.Name,
		Price: variant.Price,
	})
	if err != nil {
		return database.Variant{}, err
	}

	return insertedVariant, nil
}
