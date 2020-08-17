package services

import (
	"fmt"
	"net/url"

	"github.com/mrdulin/gqlgen-cnode/graph/model"
	"github.com/mrdulin/gqlgen-cnode/utils/http"
)

type MarkOneMessageRequestPayload struct {
	Accesstoken string `json:"accesstoken"`
}
type MarkAllMessagesRequestPayload struct {
	Accesstoken string `json:"accesstoken"`
}

type messageService struct {
	HttpClient http.Client
	BaseURL    string
}

type MessageService interface {
	GetMessages(accesstoken, mdrender string) *model.MessagesResponse
	GetUnreadMessage(accesstoken string) int
	MarkOneMessage(accesstoken, id string) *string
	MarkAllMessages(accesstoken string) []*model.MarkedMessage
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

func (m *messageService) MarkOneMessage(accesstoken, id string) *string {
	endpoint := m.BaseURL + "/message/mark_one/" + id
	var res model.MarkOneMessageResponse
	if err := m.HttpClient.Post(endpoint, &MarkOneMessageRequestPayload{Accesstoken: accesstoken}, &res); err != nil {
		fmt.Println(err)
		return &res.MarkedMsgId
	}
	return &res.MarkedMsgId
}

func (m *messageService) MarkAllMessages(accesstoken string) []*model.MarkedMessage {
	endpoint := m.BaseURL + "/message/mark_all"
	var res model.MarkAllMessagesResponse
	if err := m.HttpClient.Post(endpoint, &MarkAllMessagesRequestPayload{Accesstoken: accesstoken}, &res); err != nil {
		fmt.Println(err)
		return res.MarkedMsgs
	}
	return res.MarkedMsgs
}
