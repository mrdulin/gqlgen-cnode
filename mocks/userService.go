package mocks

import (
	"github.com/mrdulin/gqlgen-cnode/graph/model"
	"github.com/stretchr/testify/mock"
)

type MockedUserService struct {
	mock.Mock
}

func (s *MockedUserService) GetUserByLoginname(loginname string) *model.UserDetail {
	args := s.Called(loginname)
	return args.Get(0).(*model.UserDetail)
}

func (s *MockedUserService) ValidateAccessToken(accesstoken string) *model.UserEntity {
	args := s.Called(accesstoken)
	return args.Get(0).(*model.UserEntity)
}
