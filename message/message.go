package message

import (
	"fmt"
	"github.com/yulai-123/message_send_tool/lark"
	"github.com/yulai-123/twitter_message_send/conf"
	"sort"
	"time"
)

type MessageData struct {
	Author  string
	Time    time.Time
	Text    string
	ImgList []string
	TweetID string
}

var (
	larkMessage *lark.LarkMessage
)

func InitLarkMessage() {
	larkMessage = lark.NewLarkMessage(conf.GetConf().APPID, conf.GetConf().APPSecret)
}

func SendMessageWithLark(ms []MessageData) error {
	l, _ := time.LoadLocation("Asia/Shanghai")
	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Time.Before(ms[j].Time)
	})

	for _, m := range ms {
		m2 := lark.MessageData{
			Content: lark.MessageContent{
				Title:   m.Author + " 的推特消息",
				ImgList: m.ImgList,
				Content: fmt.Sprintf("发布人: %v\n发布时间: %v\n发布内容: \n%v\n",
					m.Author, m.Time.In(l).Format(time.DateTime), m.Text),
				URL: fmt.Sprintf("https://twitter.com/%v/status/%v", m.Author, m.TweetID),
			},
			ReceiveID: conf.GetConf().ReceiveID,
		}
		err := larkMessage.SendMessage(m2)
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}

	return nil

}
