package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required, max=100, min=1" json:"name"`
	Kode string `validate:"required, max=5, min=5" json:"kode"`
}
