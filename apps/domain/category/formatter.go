package category

func CategoryCreateResponse(category *Category) *Category {
	return &Category{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt:         category.CreatedAt,
	}
}
