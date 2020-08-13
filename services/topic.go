package services

import (
	"fmt"
	"net/url"

	"github.com/mrdulin/gqlgen-cnode/graph/model"
	"github.com/mrdulin/gqlgen-cnode/utils"
)

type topicService struct {
	HttpClient utils.IHttpClient
	BaseURL    string
}

type TopicService interface {
	GetTopicsByPage(urlValues *url.Values) []*model.Topic
	GetTopicById(id string) interface{}
}

func NewTopicService(httpClient utils.IHttpClient, BaseURL string) *topicService {
	return &topicService{HttpClient: httpClient, BaseURL: BaseURL}
}
func (t *topicService) GetTopicsByPage(urlValues *url.Values) []*model.Topic {
	base, err := url.Parse(t.BaseURL + "/topics")
	var res []*model.Topic
	if err != nil {
		fmt.Println("Get topics by page error: parse url.", err)
		return res
	}
	base.RawQuery = urlValues.Encode()
	body, err := t.HttpClient.Get(base.String())
	if err != nil {
		fmt.Println(err)
		return res
	}
	return body.([]*model.Topic)
}

func (t *topicService) GetTopicById(id string) interface{} {
	endpoint := t.BaseURL + "/topic/" + id
	body, err := t.HttpClient.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return body
	}
	return body
}
