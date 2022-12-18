package category

type CategoryCreateInput struct {
	Type string `json:"type" binding:"required"`
}

type CategoryUpdateInput struct {
	Type string `json:"type" binding:"required"`
}
