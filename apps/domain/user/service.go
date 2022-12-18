package user

import (
	jwt "github.com/hafizdkn/toko-belanja/apps/domain/auth/jwt"
	servicepassword "github.com/hafizdkn/toko-belanja/apps/domain/utils/service_password"
)

type IService interface {
	Create(input *UserRegisterInput) (*User, error)
	Login(input *UserLoginInput) (*UserToken, error)
	UpdateBalance(email string, input *UserTopUpInput) error
	GetUserById(input int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUsers() ([]*User, error)
	// UpdateUser(input UserUpdateInput) (UserResponse, error)
	// DeleteUser(id int) error
}

type service struct {
	repo        IRepository
	jwt         jwt.IService
	genPassword servicepassword.IService
}

func NewUserService(repository IRepository) IService {
	igenPassword := servicepassword.NewGenPassword
	ijwt := jwt.NewJwtService
	return &service{repo: repository, genPassword: igenPassword, jwt: ijwt}
}

func (s service) UpdateBalance(email string, input *UserTopUpInput) error {
	balanace := input.Balance

	err := s.repo.UpdateBalance(email, balanace)
	if err != nil {
		return err
	}

	return nil
}

func (s service) Create(input *UserRegisterInput) (*User, error) {
	var user *User

	passwordHashed, err := s.genPassword.GeneratePasswordHash(input.Password)
	if err != nil {
		return user, err
	}

	user = &User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: passwordHashed,
		Balance:  0,
		Role:     input.Role,
	}

	user, err = s.repo.CreateUser(user)
	if err != nil {
		return user, err
	}

	return RegisterResponse(user), nil
}

func (s service) Login(input *UserLoginInput) (*UserToken, error) {
	var userResponse *UserToken
	var jwtInput jwt.JwtInput

	email := input.Email
	InputPassword := input.Password

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return userResponse, err
	}

	if err := s.genPassword.ComparePasswordHash(user.Password, InputPassword); err != nil {
		return userResponse, err
	}

	jwtInput = jwt.JwtInput{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
	}

	token, err := s.jwt.GenerateToken(jwtInput)
	if err != nil {
		return userResponse, err
	}

	userResponse = &UserToken{
		Token: token,
	}

	return userResponse, nil
}

func (s service) GetUserById(id int) (*User, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s service) GetUserByEmail(email string) (*User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s service) GetUsers() ([]*User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return users, err
	}

	return users, nil
}
