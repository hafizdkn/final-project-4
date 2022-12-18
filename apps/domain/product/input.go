package product

type ProductCreateInput struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	CategoryId int    `json:"category_id" binding:"required"`
}

type ProductUpdateInput struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	CategoryId int    `json:"category_id" binding:"required"`
}
