package model

type Reply struct {
	ID       string    `json:"id"`
	Content  *string   `json:"content"`
	CreateAt *string   `json:"create_at"`
	ReplyID  *string   `json:"reply_id"`
	IsUped   *bool     `json:"is_uped"`
	Ups      []*string `json:"ups"`
	Author   *User     `json:"author"`
}

type ReplyForMessage struct {
	ID       *string   `json:"id"`
	Content  *string   `json:"content"`
	CreateAt *string   `json:"create_at"`
	Ups      []*string `json:"ups"`
}

type ReplyRecent struct {
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	LastReplyAt *string `json:"last_reply_at"`
	Author      *User   `json:"author"`
}
