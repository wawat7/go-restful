package web

type CategoryCreateRequest struct {
	Name string `validate:"required,max=255,min=1"`
}
