package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required, max=100, min=1"`
	Kode string `validate:"required, max=5, min=5"`
}