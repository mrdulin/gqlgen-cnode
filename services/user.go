package services

import (
	"fmt"

	"github.com/mrdulin/gqlgen-cnode/graph/model"
	"github.com/mrdulin/gqlgen-cnode/utils/httpClient"
)

type validateAccessTokenRequestPayload struct {
	AccessToken string `json:"accesstoken"`
}

type userService struct {
	HttpClient httpClient.HttpClient
	BaseURL    string
}

type UserService interface {
	GetUserByLoginname(loginname string) *model.UserDetail
	ValidateAccessToken(accesstoken string) *model.User
}

func NewUserService(httpClient httpClient.HttpClient, BaseURL string) *userService {
	return &userService{HttpClient: httpClient, BaseURL: BaseURL}
}

func (u *userService) GetUserByLoginname(loginname string) *model.UserDetail {
	endpoint := u.BaseURL + "/user/" + loginname
	var res *model.UserDetail
	err := u.HttpClient.Get(endpoint, &res)
	if err != nil {
		fmt.Println(err)
		return res
	}
	return res
}

func (u *userService) ValidateAccessToken(accesstoken string) *model.User {
	url := u.BaseURL + "/accesstoken"
	var res *model.User
	err := u.HttpClient.Post(url, &validateAccessTokenRequestPayload{AccessToken: accesstoken}, &res)
	if err != nil {
		fmt.Println(err)
		return res
	}
	return res
}
