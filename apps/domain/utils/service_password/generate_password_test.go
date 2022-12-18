package servicepassword

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	genPassMock = GenPasswordMock{
		Mock: mock.Mock{},
	}
	svcGenPass = NewGenPassword
)

func TestGeneratePasswordSuccess(t *testing.T) {
	password := "ini password rahasia"
	genPassMock.Mock.On("GenerateFromPassword").Return(password, nil)

	passwordHash, err := svcGenPass.GeneratePasswordHash(password)
	require.Nil(t, err)
	require.NotEmpty(t, passwordHash)
}

func TestComparePasswordHash(t *testing.T) {
	password := "ini password rahasia"
	passwordHashed := "$2a$04$unokhtQrWlHfyKvFtSeQ..6YTMXrDpSb6bqkSZjB4PKlkzGxCDHRy"

	genPassMock.Mock.On("GenerateFromPassword")

	err := svcGenPass.ComparePasswordHash(passwordHashed, password)
	require.Nil(t, err)
}
