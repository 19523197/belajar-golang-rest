package service

import (
	"belajar-golang-rest/helper"
	"belajar-golang-rest/model/domain"
	"belajar-golang-rest/model/web"
	"belajar-golang-rest/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	repository repository.CategoryRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		repository: repository,
		DB:         DB,
		Validate:   validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	category := domain.Category{
		Name: request.Name,
		Kode: request.Kode,
	}

	category = service.repository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	category, err := service.repository.FindOne(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name
	category.Kode = request.Kode

	category = service.repository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	category, err := service.repository.FindOne(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.repository.Delete(ctx, tx, category)

}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	category, err := service.repository.FindOne(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	categories, err := service.repository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToCategoryResponses(categories)
}
