package repository

import (
	"context"
	"gtmx/src/database"
)

type CatalogRepository struct {
	db *database.Queries
}

func (r CatalogRepository) ListCategories(ctx context.Context) ([]database.Category, error) {
	categories, err := r.db.ListCategories(ctx)
	if err != nil {
		return []database.Category{}, err

	}

	return categories, nil
}

func (r CatalogRepository) GetOneCategoryById(ctx context.Context, id int64) (database.Category, error) {
	product, err := r.db.GetCategory(ctx, id)

	if err != nil {
		return database.Category{}, err

	}

	return product, nil
}

func (r CatalogRepository) InsertCategory(ctx context.Context, category database.Category) (database.Category, error) {

	insertedCategory, err := r.db.CreateCategory(ctx, database.CreateCategoryParams{
		Name:      category.Name,
		SectionID: category.SectionID,
	})
	if err != nil {
		return database.Category{}, err
	}

	return insertedCategory, nil
}

func (r CatalogRepository) ListSections(ctx context.Context) ([]database.Section, error) {
	sections, err := r.db.ListSections(ctx)
	if err != nil {
		return []database.Section{}, err

	}

	return sections, nil
}

func (r CatalogRepository) GetOneSectionById(ctx context.Context, id int64) (database.Section, error) {
	section, err := r.db.GetSection(ctx, id)

	if err != nil {
		return database.Section{}, err

	}

	return section, nil
}

func (r CatalogRepository) InsertSection(ctx context.Context, section database.Section) (database.Section, error) {

	insertedSection, err := r.db.CreateSection(ctx, section.Name)
	if err != nil {
		return database.Section{}, err
	}

	return insertedSection, nil
}

func (r CatalogRepository) ListProducts(ctx context.Context) ([]database.Product, error) {
	products, err := r.db.ListProducts(ctx)
	if err != nil {
		return []database.Product{}, err

	}

	return products, nil
}

func (r CatalogRepository) GetOneProductById(ctx context.Context, id int64) (database.Product, error) {
	product, err := r.db.GetProduct(ctx, id)

	if err != nil {
		return database.Product{}, err

	}

	return product, nil
}

func (r CatalogRepository) InsertProduct(ctx context.Context, product database.Product) (database.Product, error) {

	insertedProduct, err := r.db.CreateProduct(ctx, database.CreateProductParams{
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
	variants, err := r.db.ListVariants(ctx)
	if err != nil {
		return []database.Variant{}, err

	}

	return variants, nil
}

func (r CatalogRepository) GetOneVariantById(ctx context.Context, id int64) (database.Variant, error) {
	variant, err := r.db.GetVariant(ctx, id)

	if err != nil {
		return database.Variant{}, err

	}

	return variant, nil
}

func (r CatalogRepository) InsertVariant(ctx context.Context, variant database.Variant) (database.Variant, error) {

	insertedVariant, err := r.db.CreateVariant(ctx, database.CreateVariantParams{
		Name:  variant.Name,
		Price: variant.Price,
	})
	if err != nil {
		return database.Variant{}, err
	}

	return insertedVariant, nil
}

func NewCatalogRepository(db *database.Queries) CatalogRepository {
	return CatalogRepository{db: db}
}
