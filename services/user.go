package services

import (
	"fmt"

	"github.com/mrdulin/gqlgen-cnode/utils/http"

	"github.com/mrdulin/gqlgen-cnode/graph/model"
)

type validateAccessTokenRequestPayload struct {
	AccessToken string `json:"accesstoken"`
}

type userService struct {
	HttpClient http.Client
	BaseURL    string
}

type UserService interface {
	GetUserByLoginname(loginname string) *model.UserDetail
	ValidateAccessToken(accesstoken string) *model.UserEntity
}

func NewUserService(httpClient http.Client, BaseURL string) *userService {
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

func (u *userService) ValidateAccessToken(accesstoken string) *model.UserEntity {
	url := u.BaseURL + "/accesstoken"
	var res *model.UserEntity
	err := u.HttpClient.Post(url, &validateAccessTokenRequestPayload{AccessToken: accesstoken}, &res)
	if err != nil {
		fmt.Println(err)
		return res
	}
	return res
}
