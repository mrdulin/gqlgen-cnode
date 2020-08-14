package model

type User struct {
	Loginname *string `json:"loginname"`
	AvatarURL *string `json:"avatar_url"`
}

type UserEntity struct {
	ID string `json:"id"`
	User
}

type UserDetail struct {
	User
	GithubUsername *string        `json:"githubUsername"`
	CreateAt       *string        `json:"create_at"`
	Score          *int           `json:"score"`
	RecentReplies  []*ReplyRecent `json:"recent_replies"`
	RecentTopics   []*TopicRecent `json:"recent_topics"`
}
