package repository

import (
	"belajar-golang-rest/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error)
	FindOne(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error)
	Save(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, Category domain.Category)
}
