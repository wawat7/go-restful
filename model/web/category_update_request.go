package web

type CategoryUpdateRequest struct {
	Id int `validate:"required"`
	Name string `validate:"required,max=255,min=1"`
}
