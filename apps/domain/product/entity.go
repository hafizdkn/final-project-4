package product

import "time"

type Product struct {
	ID         int        `json:"id,omitempty" gorm:"primaryKey"`
	Title      string     `json:"title,omitempty" gorm:"not null"`
	Price      int        `json:"price,omitempty" gorm:"not null"`
	Stock      int        `json:"stock,omitempty" gorm:"not null"`
	CategoryId int        `json:"category_id,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}
