package message

import "time"

type MessageData struct {
	Author  string
	Time    time.Time
	Text    string
	ImgList []string
	TweetID string
}
