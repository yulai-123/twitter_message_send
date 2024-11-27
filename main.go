package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yulai-123/twitter_message_send/conf"
	"github.com/yulai-123/twitter_message_send/message"
	"github.com/yulai-123/twitter_message_send/twitter"
	"strings"
	"time"
)

func main() {
	conf.InitConfig("conf/my_config.yaml")
	message.InitLarkMessage()
	t := time.NewTicker(5 * time.Second)
	messageLock := make(map[string]int64)
	l, _ := time.LoadLocation("Asia/Shanghai")

	for {
		select {
		case <-t.C:
			logrus.Info("触发定时器", time.Now().In(l).Format(time.DateTime))

			messageList := getMessageList(messageLock)
			err := message.SendMessageWithLark(messageList)
			if err != nil {
				logrus.Info("触发定时器", time.Now().In(l).Format(time.DateTime))
				continue
			}
			for _, m := range messageList {
				messageLock[m.TweetID] = time.Now().Unix()
			}

			for key, lockTime := range messageLock {
				if time.Now().Add(-2 * 24 * time.Hour).After(time.Unix(lockTime, 0)) {
					delete(messageLock, key)
				}
			}
		}
	}
}

func getMessageList(messageLock map[string]int64) []message.MessageData {
	messageList := make([]message.MessageData, 0)
	for userID, author := range conf.GetConf().UserIDs {
		postList, err := twitter.GetPostList(userID)
		if err != nil {
			s := fmt.Sprintf("拉取帖子列表失败, userID: %v, err: %v\n", userID, err)
			messageList = appendError(messageList, fmt.Sprintf("发生一个错误, %v", s))
			continue
		}

		instructions := postList.Data.User.Result.TimelineV2.Timeline.Instructions

		for _, ins := range instructions {
			if ins.Type != "TimelineAddEntries" && ins.Type != "TimelinePinEntry" {
				continue
			}

			entries := ins.Entries
			for _, entry := range entries {
				tweetId := entry.Content.ItemContent.TweetResults.Result.RestID
				if tweetId != "" {
					createdAt, err := time.Parse(time.RubyDate, entry.Content.ItemContent.TweetResults.Result.Legacy.CreatedAt)
					if err != nil {
						s := fmt.Sprintf("解析时间错误, tweetId: %v, createdAt: %v\n", tweetId, entry.Content.ItemContent.TweetResults.Result.Legacy.CreatedAt)
						messageList = appendError(messageList, fmt.Sprintf("发生一个错误, %v", s))
						break
					}

					if time.Now().Add(-1.5 * 24 * time.Hour).After(createdAt) {
						continue
					}

					if _, exist := messageLock[tweetId]; exist {
						continue
					}

					post, err := twitter.GetPostV2(tweetId)
					if err != nil {
						s := fmt.Sprintf("获取post失败，err: %v\n", err)
						messageList = appendError(messageList, fmt.Sprintf("发生一个错误, %v", s))
						break
					}

					for _, inst := range post.Data.ThreadedConversationWithInjectionsV2.Instructions {
						for _, ent := range inst.Entries {
							if ent.Content.ItemContent.TweetResults.Result.RestID == tweetId {
								text := ent.Content.ItemContent.TweetResults.Result.NoteTweet.NoteTweetResults.Result.Text
								if len(text) < 5 {
									text = ent.Content.ItemContent.TweetResults.Result.Legacy.FullText
								}
								messageList = append(messageList, message.MessageData{
									Author:  author,
									Time:    createdAt,
									Text:    text,
									ImgList: nil,
									TweetID: tweetId,
								})
								continue
							}

							for _, ite := range ent.Content.Items {
								if strings.Contains(ite.EntryID, tweetId) {
									text := ent.Content.ItemContent.TweetResults.Result.NoteTweet.NoteTweetResults.Result.Text
									if len(text) < 5 {
										text = ent.Content.ItemContent.TweetResults.Result.Legacy.FullText
									}
									messageList = append(messageList, message.MessageData{
										Author:  author,
										Time:    createdAt,
										Text:    text,
										ImgList: nil,
										TweetID: tweetId,
									})
									continue
								}
							}
						}
					}

					continue
				}

				for _, item := range entry.Content.Items {
					tweetId := item.Item.ItemContent.TweetResults.Result.RestID
					if tweetId == "" {
						continue
					}
					createdAt, err := time.Parse(time.RubyDate, item.Item.ItemContent.TweetResults.Result.Legacy.CreatedAt)
					if err != nil {
						s := fmt.Sprintf("获取post失败，err: %v\n", err)
						messageList = appendError(messageList, fmt.Sprintf("发生一个错误, %v", s))
						break
					}

					if time.Now().Add(-1.5 * 24 * time.Hour).After(createdAt) {
						continue
					}
					if _, exist := messageLock[tweetId]; exist {
						continue
					}

					post, err := twitter.GetPostV2(tweetId)
					if err != nil {
						s := fmt.Sprintf("获取post失败，err: %v\n", err)
						messageList = appendError(messageList, fmt.Sprintf("发生一个错误, %v", s))
						break
					}

					for _, inst := range post.Data.ThreadedConversationWithInjectionsV2.Instructions {
						for _, ent := range inst.Entries {
							if ent.Content.ItemContent.TweetResults.Result.RestID == tweetId {
								text := ent.Content.ItemContent.TweetResults.Result.NoteTweet.NoteTweetResults.Result.Text
								if len(text) < 5 {
									text = ent.Content.ItemContent.TweetResults.Result.Legacy.FullText
								}
								messageList = append(messageList, message.MessageData{
									Author:  author,
									Time:    createdAt,
									Text:    text,
									ImgList: nil,
									TweetID: tweetId,
								})
								continue
							}

							for _, ite := range ent.Content.Items {
								if strings.Contains(ite.EntryID, tweetId) {
									text := ent.Content.ItemContent.TweetResults.Result.NoteTweet.NoteTweetResults.Result.Text
									if len(text) < 5 {
										text = ent.Content.ItemContent.TweetResults.Result.Legacy.FullText
									}
									messageList = append(messageList, message.MessageData{
										Author:  author,
										Time:    createdAt,
										Text:    text,
										ImgList: nil,
										TweetID: tweetId,
									})
									continue
								}
							}
						}
					}
				}

			}

		}
	}

	return messageList
}

var (
	lastErrorTime time.Time = time.Now()
)

func appendError(messageList []message.MessageData, errStr string) []message.MessageData {
	// 错误信息十分钟间隔发送
	if time.Now().After(lastErrorTime.Add(10 * time.Minute)) {
		messageList = append(messageList, message.MessageData{
			Author: "error",
			Time:   time.Now(),
			Text:   errStr,
		})

		lastErrorTime = time.Now()
	}

	if !time.Now().After(lastErrorTime.Add(10 * time.Minute)) {
		logrus.Errorf("跳过错误: %v", errStr)
	}

	return messageList
}
