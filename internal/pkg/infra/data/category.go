package data

import (
	"context"
	"fmt"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	pkgRepo "github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type CategoryRepository struct {
	conn         pkgRepo.Connector
	mainTable    string
	selectFields string
}

func NewCategoryRepository(conn pkgRepo.Connector) repository.CategoryRepository {
	return CategoryRepository{
		conn:      conn,
		mainTable: "tb_category",
		selectFields: `
			id
			, created_at
			, updated_at
			, slug
			, name
			
			FROM tb_category
		`,
	}
}

func (r CategoryRepository) parseEntity(s pkgRepo.Scanner) (entity.Category, error) {
	var slug string
	var category entity.Category
	err := s.Scan(
		&category.ID,
		&category.CreatedAt,
		&category.UpdatedAt,
		&slug,
		&category.Name,
	)
	if err != nil {
		return entity.Category{}, errors.Repo(err, r.mainTable)
	}
	category.SetSlug(slug)

	return category, nil
}

func (r CategoryRepository) Insert(ctx context.Context, category entity.Category) (uint32, error) {
	res, err := r.conn.ExecContext(ctx,
		"INSERT INTO tb_category (slug, name) VALUES (?, ?)",
		category.GetSlug(),
		category.Name,
	)
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	return uint32(id), nil
}

func (r CategoryRepository) FindBySlug(ctx context.Context, slug string) (entity.Category, error) {
	query := `
		SELECT %s
			WHERE NOT deleted_at AND slug = ?
	`

	q := fmt.Sprintf(query, r.selectFields)

	category, err := r.parseEntity(
		r.conn.QueryRowContext(
			ctx,
			q,
			slug,
		))
	if err != nil {
		return entity.Category{}, errors.Repo(err, r.mainTable)
	}

	return category, nil
}

func (r CategoryRepository) FindAll(ctx context.Context) ([]entity.Category, error) {
	query := `
		SELECT %s
			WHERE NOT deleted_at
	`

	q := fmt.Sprintf(query, r.selectFields)

	result, err := r.conn.QueryContext(
		ctx,
		q,
	)

	list, err := pkgRepo.ParseEntities[entity.Category](
		r.parseEntity,
		result,
		err,
	)
	if err != nil {
		return []entity.Category{}, errors.Repo(err, r.mainTable)
	}

	return list, nil
}
