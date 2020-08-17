package services_test

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/mrdulin/gqlgen-cnode/graph/model"

	"github.com/mrdulin/gqlgen-cnode/services"

	"github.com/mrdulin/gqlgen-cnode/mocks"
)

const (
	baseURL     string = "http://localhost/api/v1"
	accesstoken string = "123"
)

func TestMessageService_MarkAllMessages(t *testing.T) {
	t.Run("should mark all messaages", func(t *testing.T) {
		testHttp := new(mocks.MockedHttp)
		var res model.MarkAllMessagesResponse
		var markedMsgs []*model.MarkedMessage
		for i := 1; i <= 3; i++ {
			markedMsgs = append(markedMsgs, &model.MarkedMessage{ID: strconv.Itoa(i)})
		}
		postBody := services.MarkAllMessagesRequestPayload{Accesstoken: accesstoken}
		testHttp.On("Post", baseURL+"/message/mark_all", &postBody, &res).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(2).(*model.MarkAllMessagesResponse)
			arg.MarkedMsgs = markedMsgs
		})
		service := services.NewMessageService(testHttp, baseURL)
		got := service.MarkAllMessages(accesstoken)
		want := markedMsgs
		testHttp.AssertExpectations(t)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got wrong return value. got: %#v, want: %#v", got, want)
		}
	})

	t.Run("should print error and return empty slice", func(t *testing.T) {
		var res model.MarkAllMessagesResponse
		testHttp := new(mocks.MockedHttp)
		postBody := services.MarkAllMessagesRequestPayload{Accesstoken: accesstoken}
		testHttp.On("Post", baseURL+"/message/mark_all", &postBody, &res).Return(errors.New("network"))
		service := services.NewMessageService(testHttp, baseURL)
		got := service.MarkAllMessages(accesstoken)
		var want []*model.MarkedMessage
		testHttp.AssertExpectations(t)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got wrong return value. got: %#v, want: %#v", got, want)
		}
	})
}
