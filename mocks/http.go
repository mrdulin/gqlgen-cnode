package mocks

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockedHttp struct {
	mock.Mock
}

func (m *MockedHttp) Get(url string, data interface{}) error {
	args := m.Called(url, data)
	return args.Error(0)
}

func (m *MockedHttp) Post(url string, body interface{}, data interface{}) error {
	args := m.Called(url, body, data)
	return args.Error(0)
}

func (m *MockedHttp) HandleAPIError(res interface{}) error {
	args := m.Called(res)
	return args.Error(0)
}

func (m *MockedHttp) Decode(body io.ReadCloser, res interface{}) error {
	args := m.Called(body, res)
	return args.Error(0)
}
func (m *MockedHttp) Unmarshal(byte interface{}, data interface{}) error {
	args := m.Called(byte, data)
	return args.Error(0)
}
