package request

type CategoryRequest struct {
	CategoryName string `json:"category_name" binding:"required,min=3,max=50"`
}
