package productresponse

import "time"

type ProductResponse struct {
	ID         int        `json:"id,omitempty"`
	Title      string     `json:"title,omitempty"`
	Price      int        `json:"price,omitempty"`
	Stock      int        `json:"stock,omitempty"`
	CategoryId int        `json:"-"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

func (ProductResponse) TableName() string {
	return "products"
}
