package services

import (
	"fmt"
	"net/url"

	"github.com/mrdulin/gqlgen-cnode/graph/model"
	"github.com/mrdulin/gqlgen-cnode/utils/http"
)

type messageService struct {
	HttpClient http.Client
	BaseURL    string
}

type MessageService interface {
	GetMessages(accesstoken, mdrender string) *model.MessagesResponse
	GetUnreadMessage(accesstoken string) int
}

func NewMessageService(httpClient http.Client, BaseURL string) *messageService {
	return &messageService{HttpClient: httpClient, BaseURL: BaseURL}
}

func (m *messageService) GetMessages(accesstoken, mdrender string) *model.MessagesResponse {
	base, err := url.Parse(m.BaseURL + "/messages")
	var res *model.MessagesResponse
	if err != nil {
		fmt.Println("Get messages parse url error.", err)
		return res
	}
	urlValues := url.Values{}
	urlValues.Add("accesstoken", accesstoken)
	urlValues.Add("mdrender", mdrender)
	base.RawQuery = urlValues.Encode()
	if err = m.HttpClient.Get(base.String(), &res); err != nil {
		fmt.Println(err)
		return res
	}
	return res
}

func (m *messageService) GetUnreadMessage(accesstoken string) int {
	base, err := url.Parse(m.BaseURL + "/message/count")
	var res int
	if err != nil {
		fmt.Println("Get unread message parse url error.", err)
		return res
	}
	urlValues := url.Values{}
	urlValues.Add("accesstoken", accesstoken)
	base.RawQuery = urlValues.Encode()
	if err = m.HttpClient.Get(base.String(), &res); err != nil {
		fmt.Println(err)
		return res
	}
	return res
}
