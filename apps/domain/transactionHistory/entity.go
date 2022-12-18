package transactionhistory

import (
	"time"

	productresponse "github.com/hafizdkn/toko-belanja/apps/domain/utils/product_response"
	userresponse "github.com/hafizdkn/toko-belanja/apps/domain/utils/user_response"
)

type Transaction struct {
	ID         int                              `json:"id,omitempty" gorm:"primaryKey"`
	ProductID  int                              `json:"product_id,omitempty"`
	UserId     int                              `json:"user_id,omitempty"`
	Quantity   int                              `json:"quantity,omitempty" gorm:"not null"`
	TotalPrice int                              `json:"total_price,omitempty" gorm:"not null"`
	CreatedAt  *time.Time                       `json:"created_at,omitempty"`
	UpdatedAt  *time.Time                       `json:"updated_at,omitempty"`
	Product    *productresponse.ProductResponse `json:"product,omitempty"`
	User       userresponse.UserResponse        `json:"user"`
}
