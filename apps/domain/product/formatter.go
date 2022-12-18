package product

func ProductResponseRequest(product *Product) *Product {
	return &Product{
		ID:         product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryId: product.CategoryId,
		CreatedAt:  product.CreatedAt,
	}
}
