package web

type BookUpdateRequest struct {
	Id    int    `validate:"required"`
	Title string `validate:"required,max=255,min=1" json:"title"`
}
