package repository

import (
	"belajar-golang-rest/helper"
	"belajar-golang-rest/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	SQL := "INSERT INTO category (name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category set name = ?, kode = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Kode, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error) {
	SQL := "SELECT id, nama, kode FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	categories := []domain.Category{}
	category := domain.Category{}
	for rows.Next() {
		err := rows.Scan(&category.Id, &category.Name, &category.Kode)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories, nil
}

func (repository *CategoryRepositoryImpl) FindOne(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	SQL := "SELECT id, nama, kode FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name, &category.Kode)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Category not found")
	}
}
