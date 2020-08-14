package services

import (
	"fmt"
	"net/url"

	"github.com/mrdulin/gqlgen-cnode/utils/http"

	"github.com/google/go-querystring/query"

	"github.com/mrdulin/gqlgen-cnode/graph/model"
)

type topicService struct {
	HttpClient http.Client
	BaseURL    string
}

type TopicService interface {
	GetTopicsByPage(params *model.TopicsRequestParams) []*model.Topic
	GetTopicById(params *model.TopicRequestParams) *model.TopicDetail
}

func NewTopicService(httpClient http.Client, BaseURL string) *topicService {
	return &topicService{HttpClient: httpClient, BaseURL: BaseURL}
}
func (t *topicService) GetTopicsByPage(params *model.TopicsRequestParams) []*model.Topic {
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

func (t *topicService) GetTopicById(params *model.TopicRequestParams) *model.TopicDetail {
	base, err := url.Parse(t.BaseURL + "/topic/" + params.ID)
	var res *model.TopicDetail
	if err != nil {
		fmt.Println("Get topic by id error: parse url.", err)
		return res
	}
	urlValues := url.Values{}
	if params.Accesstoken != nil {
		urlValues.Add("accesstoken", *params.Accesstoken)
	}
	urlValues.Add("mdrender", *params.Mdrender)
	base.RawQuery = urlValues.Encode()
	err = t.HttpClient.Get(base.String(), &res)
	if err != nil {
		fmt.Println(err)
		return res
	}
	return res
}
