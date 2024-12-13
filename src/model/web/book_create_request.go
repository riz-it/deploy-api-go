package web

type BookCreateRequest struct {
	Title string `validate:"required,max=255,min=1" json:"title"`
}
