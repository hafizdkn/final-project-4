package userresponse

import "time"

type UserResponse struct {
	ID        int        `json:"id,omitempty"`
	Email     string     `json:"email,omitempty"`
	FullName  string     `json:"full_name,omitempty"`
	Balance   string     `json:"balance,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (UserResponse) TableName() string {
	return "users"
}
