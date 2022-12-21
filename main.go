package main

import (
	"belajar-golang-rest/app"
	"belajar-golang-rest/controller"
	"belajar-golang-rest/exception"
	"belajar-golang-rest/helper"
	"belajar-golang-rest/middleware"
	"belajar-golang-rest/model/web"
	"belajar-golang-rest/repository"
	"belajar-golang-rest/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api", Welcome)
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories:categoryId", categoryController.Update)
	router.DELETE("/api/categories:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:8001",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

func Welcome(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   "Server Up and Running",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
