package user

import "time"

type User struct {
	ID        int        `json:"id,omitempty" gorm:"primaryKey"`
	FullName  string     `json:"full_name,omitempty" gorm:"not null"`
	Email     string     `json:"email,omitempty" gorm:"not null;unique"`
	Password  string     `json:"password,omitempty" gorm:"not null"`
	Role      string     `json:"role,omitempty" gorm:"not null"`
	Balance   int        `json:"balance,omitempty" gorm:"not null"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UserToken struct {
	Token string `json:"token,omitempty"`
}
