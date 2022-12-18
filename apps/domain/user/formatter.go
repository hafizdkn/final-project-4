package user

func RegisterResponse(user *User) *User {
	return &User{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Balance:   0,
		CreatedAt: user.CreatedAt,
	}
}
