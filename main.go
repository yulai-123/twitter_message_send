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
	message.InitLarkMessage()
	t := time.NewTicker(time.Minute)
	messageLock := make(map[string]int64)

	for {
		select {
		case <-t.C:
			fmt.Println("触发定时器")
			messageList := getMessageList(messageLock)
			err := message.SendMessageWithLark(messageList)
			if err != nil {
				fmt.Println(err)
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
			fmt.Println(s)
			messageList = append(messageList, message.MessageData{
				Author: "error",
				Time:   time.Now(),
				Text:   fmt.Sprintf("发生一个错误, %v", s),
			})
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
						s := fmt.Sprintf("获取post失败，err: %v\n", err)
						fmt.Println(s)
						messageList = append(messageList, message.MessageData{
							Author: "error",
							Time:   time.Now(),
							Text:   fmt.Sprintf("发生一个错误, %v", s),
						})
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
						s := fmt.Sprintf("获取post失败，err: %v\n", err)
						fmt.Println(s)
						messageList = append(messageList, message.MessageData{
							Author: "error",
							Time:   time.Now(),
							Text:   fmt.Sprintf("发生一个错误, %v", s),
						})
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
