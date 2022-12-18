package servicepassword

import (
	"github.com/stretchr/testify/mock"
)

type GenPasswordMock struct {
	Mock mock.Mock
}

func (g *GenPasswordMock) GenerateFromPassword(password string) (string, error) {
	args := g.Mock.Called(password)

	if args.Get(0) != nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), nil
}

func (g *GenPasswordMock) ComparePasswordHash(passwordHashed, passwrod string) error {
	args := g.Mock.Called(passwordHashed, passwrod)

	if args.Get(1) != nil {
		return args.Error(1)
	}

	return args.Get(1).(error)
}
