package model

type Message struct {
	ID       string           `json:"id"`
	Type     *string          `json:"type"`
	HasRead  *bool            `json:"has_read"`
	CreateAt *string          `json:"create_at"`
	Reply    *ReplyForMessage `json:"reply"`
	Topic    *TopicForMessage `json:"topic"`
	Author   *User            `json:"author"`
}

type MessagesResponse struct {
	HasReadMessages    []*Message `json:"has_read_messages"`
	HasnotReadMessages []*Message `json:"hasnot_read_messages"`
}
