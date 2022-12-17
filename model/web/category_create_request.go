package web

type CategoryCreateRequest struct {
	Name string `validate:"required, max=100, min=1"`
	Kode string `validate:"required, max=5, min=5"`
}
