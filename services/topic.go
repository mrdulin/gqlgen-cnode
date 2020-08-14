package services

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/mrdulin/gqlgen-cnode/graph/model"
	"github.com/mrdulin/gqlgen-cnode/utils/httpClient"
)

type topicService struct {
	HttpClient httpClient.HttpClient
	BaseURL    string
}

type TopicService interface {
	GetTopicsByPage(params model.TopicsRequestParams) []*model.Topic
	GetTopicById(id string) *model.TopicDetail
}

func NewTopicService(httpClient httpClient.HttpClient, BaseURL string) *topicService {
	return &topicService{HttpClient: httpClient, BaseURL: BaseURL}
}
func (t *topicService) GetTopicsByPage(params model.TopicsRequestParams) []*model.Topic {
	base, err := url.Parse(t.BaseURL + "/topics")
	var res []*model.Topic
	if err != nil {
		fmt.Println("Get topics by page error: parse url.", err)
		return res
	}
	v, err := query.Values(params)
	if err != nil {
		fmt.Printf("query.Values(params) error. params: %+v, error: %s", params, err)
		return res
	}
	base.RawQuery = v.Encode()
	err = t.HttpClient.Get(base.String(), &res)
	if err != nil {
		fmt.Println(err)
		return res
	}
	return res
}

func (t *topicService) GetTopicById(id string) *model.TopicDetail {
	endpoint := t.BaseURL + "/topic/" + id
	res := model.TopicDetail{}
	err := t.HttpClient.Get(endpoint, &res)
	if err != nil {
		fmt.Println(err)
		return &res
	}
	return &res
}
