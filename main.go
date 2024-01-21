package main

import (
	"fmt"
	"github.com/yulai-123/twitter_message_send/conf"
	"github.com/yulai-123/twitter_message_send/message"
	"github.com/yulai-123/twitter_message_send/twitter"
	"strings"
	"time"
)

func main() {
	conf.InitConfig("conf/my_config.yaml")
	t := time.NewTicker(time.Minute)
	messageLock := make(map[string]int64)

	for {
		select {
		case <-t.C:
			fmt.Println("触发定时器")
			messageList := getMessageList(messageLock)

			l, _ := time.LoadLocation("Asia/Shanghai")

			for _, message := range messageList {
				fmt.Printf("发布人: %v\n发布时间: %v\n发布内容: %v\n", message.Author, message.Time.In(l).Format(time.DateTime), message.Text)
				messageLock[message.TweetID] = time.Now().Unix()
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
			fmt.Printf("拉取帖子列表失败, userID: %v, err: %v\n", userID, err)
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
						fmt.Printf("解析时间错误, tweetId: %v, createdAt: %v\n", tweetId, entry.Content.ItemContent.TweetResults.Result.Legacy.CreatedAt)
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
						fmt.Printf("获取post失败，err: %v\n", err)
						break
					}

					for _, inst := range post.Data.ThreadedConversationWithInjectionsV2.Instructions {
						for _, ent := range inst.Entries {
							if ent.Content.ItemContent.TweetResults.Result.RestID == tweetId {
								messageList = append(messageList, message.MessageData{
									Author:  author,
									Time:    createdAt,
									Text:    ent.Content.ItemContent.TweetResults.Result.Legacy.FullText,
									ImgList: nil,
									TweetID: tweetId,
								})
								continue
							}

							for _, ite := range ent.Content.Items {
								if strings.Contains(ite.EntryID, tweetId) {
									messageList = append(messageList, message.MessageData{
										Author:  author,
										Time:    createdAt,
										Text:    ite.Item.ItemContent.TweetResults.Result.Legacy.FullText,
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
						fmt.Printf("解析时间错误, tweetId: %v, createdAt: %v\n", tweetId, item.Item.ItemContent.TweetResults.Result.Legacy.CreatedAt)
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
						fmt.Printf("获取post失败，err: %v\n", err)
						break
					}

					for _, inst := range post.Data.ThreadedConversationWithInjectionsV2.Instructions {
						for _, ent := range inst.Entries {
							if ent.Content.ItemContent.TweetResults.Result.RestID == tweetId {
								messageList = append(messageList, message.MessageData{
									Author:  author,
									Time:    createdAt,
									Text:    ent.Content.ItemContent.TweetResults.Result.Legacy.FullText,
									ImgList: nil,
									TweetID: tweetId,
								})
								continue
							}

							for _, ite := range ent.Content.Items {
								if strings.Contains(ite.EntryID, tweetId) {
									messageList = append(messageList, message.MessageData{
										Author:  author,
										Time:    createdAt,
										Text:    ite.Item.ItemContent.TweetResults.Result.Legacy.FullText,
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
