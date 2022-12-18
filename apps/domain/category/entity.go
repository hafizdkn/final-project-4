package category

import (
	"time"

	productresponse "github.com/hafizdkn/toko-belanja/apps/domain/utils/product_response"
)

type Category struct {
	ID                int                              `json:"id,omitempty" gorm:"primaryKey"`
	Type              string                           `json:"type,omitempty" gorm:"not null"`
	SoldProductAmount int                              `json:"sold_product_amount,omitempty"`
	CreatedAt         *time.Time                       `json:"created_at,omitempty"`
	UpdatedAt         *time.Time                       `json:"updated_at,omitempty"`
	Products          *productresponse.ProductResponse `json:"Products,omitempty"`
}
