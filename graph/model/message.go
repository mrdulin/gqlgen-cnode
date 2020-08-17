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

type MarkOneMessageResponse struct {
	MarkedMsgId string `json:"marked_msg_id"`
}

type MarkedMessage struct {
	ID string `json:"id"`
}

type MarkAllMessagesResponse struct {
	MarkedMsgs []*MarkedMessage `json:"marked_msgs"`
}
