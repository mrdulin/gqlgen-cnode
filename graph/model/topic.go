package model

import (
	"fmt"
	"io"
	"strconv"
)

type Topic struct {
	ID          string  `json:"id"`
	AuthorID    string  `json:"author_id"`
	Tab         *string `json:"tab"`
	Content     *string `json:"content"`
	Title       string  `json:"title"`
	LastReplyAt *string `json:"last_reply_at"`
	Good        *bool   `json:"good"`
	Top         *bool   `json:"top"`
	ReplyCount  *int    `json:"reply_count"`
	VisitCount  *int    `json:"visit_count"`
	CreateAt    *string `json:"create_at"`
	IsCollect   *bool   `json:"is_collect"`
	Author      *User   `json:"author"`
}

type TopicDetail struct {
	ID          string   `json:"id"`
	AuthorID    string   `json:"author_id"`
	Tab         *string  `json:"tab"`
	Content     *string  `json:"content"`
	Title       string   `json:"title"`
	LastReplyAt *string  `json:"last_reply_at"`
	Good        *bool    `json:"good"`
	Top         *bool    `json:"top"`
	ReplyCount  *int     `json:"reply_count"`
	VisitCount  *int     `json:"visit_count"`
	CreateAt    *string  `json:"create_at"`
	IsCollect   *bool    `json:"is_collect"`
	Replies     []*Reply `json:"replies"`
	Author      *User    `json:"author"`
}

type TopicForMessage struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	LastReplyAt *string `json:"last_reply_at"`
}

type TopicRecent struct {
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	LastReplyAt *string `json:"last_reply_at"`
	Author      *User   `json:"author"`
}

type TopicRequestParams struct {
	ID          string  `json:"id" url:"id,omitempty"`
	Accesstoken *string `json:"accesstoken" url:"accesstoken,omitempty"`
	Mdrender    *string `json:"mdrender" url:"mdrender,omitempty"`
}

type TopicsRequestParams struct {
	Page     *int      `json:"page" url:"page,omitempty"`
	Tab      *TopicTab `json:"tab" url:"tab,omitempty"`
	Limit    *int      `json:"limit" url:"limit,omitempty"`
	Mdrender *string   `json:"mdrender" url:"mdrender,omitempty"`
}

type TopicTab string

const (
	TopicTabAsk   TopicTab = "ask"
	TopicTabShare TopicTab = "share"
	TopicTabJob   TopicTab = "job"
	TopicTabGood  TopicTab = "good"
)

var AllTopicTab = []TopicTab{
	TopicTabAsk,
	TopicTabShare,
	TopicTabJob,
	TopicTabGood,
}

func (e TopicTab) IsValid() bool {
	switch e {
	case TopicTabAsk, TopicTabShare, TopicTabJob, TopicTabGood:
		return true
	}
	return false
}

func (e TopicTab) String() string {
	return string(e)
}

func (e *TopicTab) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TopicTab(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TopicTab", str)
	}
	return nil
}

func (e TopicTab) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
