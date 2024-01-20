package twitter

type PostResp struct {
	Data struct {
		TweetResult struct {
			Result struct {
				Typename string `json:"__typename"`
				RestID   string `json:"rest_id"`
				Core     struct {
					UserResults struct {
						Result struct {
							Typename                   string `json:"__typename"`
							ID                         string `json:"id"`
							RestID                     string `json:"rest_id"`
							AffiliatesHighlightedLabel struct {
							} `json:"affiliates_highlighted_label"`
							IsBlueVerified    bool   `json:"is_blue_verified"`
							ProfileImageShape string `json:"profile_image_shape"`
							Legacy            struct {
								CreatedAt           string `json:"created_at"`
								DefaultProfile      bool   `json:"default_profile"`
								DefaultProfileImage bool   `json:"default_profile_image"`
								Description         string `json:"description"`
								Entities            struct {
									Description struct {
										Urls []interface{} `json:"urls"`
									} `json:"description"`
									URL struct {
										Urls []struct {
											DisplayURL  string `json:"display_url"`
											ExpandedURL string `json:"expanded_url"`
											URL         string `json:"url"`
											Indices     []int  `json:"indices"`
										} `json:"urls"`
									} `json:"url"`
								} `json:"entities"`
								FastFollowersCount      int           `json:"fast_followers_count"`
								FavouritesCount         int           `json:"favourites_count"`
								FollowersCount          int           `json:"followers_count"`
								FriendsCount            int           `json:"friends_count"`
								HasCustomTimelines      bool          `json:"has_custom_timelines"`
								IsTranslator            bool          `json:"is_translator"`
								ListedCount             int           `json:"listed_count"`
								Location                string        `json:"location"`
								MediaCount              int           `json:"media_count"`
								Name                    string        `json:"name"`
								NormalFollowersCount    int           `json:"normal_followers_count"`
								PinnedTweetIdsStr       []interface{} `json:"pinned_tweet_ids_str"`
								PossiblySensitive       bool          `json:"possibly_sensitive"`
								ProfileBannerURL        string        `json:"profile_banner_url"`
								ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
								ProfileInterstitialType string        `json:"profile_interstitial_type"`
								ScreenName              string        `json:"screen_name"`
								StatusesCount           int           `json:"statuses_count"`
								TranslatorType          string        `json:"translator_type"`
								URL                     string        `json:"url"`
								Verified                bool          `json:"verified"`
								WithheldInCountries     []interface{} `json:"withheld_in_countries"`
							} `json:"legacy"`
							Professional struct {
								RestID           string `json:"rest_id"`
								ProfessionalType string `json:"professional_type"`
								Category         []struct {
									ID       int    `json:"id"`
									Name     string `json:"name"`
									IconName string `json:"icon_name"`
								} `json:"category"`
							} `json:"professional"`
						} `json:"result"`
					} `json:"user_results"`
				} `json:"core"`
				UnmentionData struct {
				} `json:"unmention_data"`
				EditControl struct {
					EditTweetIds       []string `json:"edit_tweet_ids"`
					EditableUntilMsecs string   `json:"editable_until_msecs"`
					IsEditEligible     bool     `json:"is_edit_eligible"`
					EditsRemaining     string   `json:"edits_remaining"`
				} `json:"edit_control"`
				IsTranslatable bool `json:"is_translatable"`
				Views          struct {
					State string `json:"state"`
				} `json:"views"`
				Source string `json:"source"`
				Legacy struct {
					BookmarkCount     int    `json:"bookmark_count"`
					Bookmarked        bool   `json:"bookmarked"`
					CreatedAt         string `json:"created_at"`
					ConversationIDStr string `json:"conversation_id_str"`
					DisplayTextRange  []int  `json:"display_text_range"`
					Entities          struct {
						Hashtags []interface{} `json:"hashtags"`
						Media    []struct {
							DisplayURL           string `json:"display_url"`
							ExpandedURL          string `json:"expanded_url"`
							IDStr                string `json:"id_str"`
							Indices              []int  `json:"indices"`
							MediaKey             string `json:"media_key"`
							MediaURLHTTPS        string `json:"media_url_https"`
							Type                 string `json:"type"`
							URL                  string `json:"url"`
							ExtMediaAvailability struct {
								Status string `json:"status"`
							} `json:"ext_media_availability"`
							Features struct {
								Large struct {
									Faces []interface{} `json:"faces"`
								} `json:"large"`
								Medium struct {
									Faces []interface{} `json:"faces"`
								} `json:"medium"`
								Small struct {
									Faces []interface{} `json:"faces"`
								} `json:"small"`
								Orig struct {
									Faces []interface{} `json:"faces"`
								} `json:"orig"`
							} `json:"features"`
							Sizes struct {
								Large struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"large"`
								Medium struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"medium"`
								Small struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"small"`
								Thumb struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"thumb"`
							} `json:"sizes"`
							OriginalInfo struct {
								Height     int `json:"height"`
								Width      int `json:"width"`
								FocusRects []struct {
									X int `json:"x"`
									Y int `json:"y"`
									W int `json:"w"`
									H int `json:"h"`
								} `json:"focus_rects"`
							} `json:"original_info"`
						} `json:"media"`
						Symbols      []interface{} `json:"symbols"`
						Timestamps   []interface{} `json:"timestamps"`
						Urls         []interface{} `json:"urls"`
						UserMentions []struct {
							IDStr      string `json:"id_str"`
							Name       string `json:"name"`
							ScreenName string `json:"screen_name"`
							Indices    []int  `json:"indices"`
						} `json:"user_mentions"`
					} `json:"entities"`
					ExtendedEntities struct {
						Media []struct {
							DisplayURL           string `json:"display_url"`
							ExpandedURL          string `json:"expanded_url"`
							IDStr                string `json:"id_str"`
							Indices              []int  `json:"indices"`
							MediaKey             string `json:"media_key"`
							MediaURLHTTPS        string `json:"media_url_https"`
							Type                 string `json:"type"`
							URL                  string `json:"url"`
							ExtMediaAvailability struct {
								Status string `json:"status"`
							} `json:"ext_media_availability"`
							Features struct {
								Large struct {
									Faces []interface{} `json:"faces"`
								} `json:"large"`
								Medium struct {
									Faces []interface{} `json:"faces"`
								} `json:"medium"`
								Small struct {
									Faces []interface{} `json:"faces"`
								} `json:"small"`
								Orig struct {
									Faces []interface{} `json:"faces"`
								} `json:"orig"`
							} `json:"features"`
							Sizes struct {
								Large struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"large"`
								Medium struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"medium"`
								Small struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"small"`
								Thumb struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"thumb"`
							} `json:"sizes"`
							OriginalInfo struct {
								Height     int `json:"height"`
								Width      int `json:"width"`
								FocusRects []struct {
									X int `json:"x"`
									Y int `json:"y"`
									W int `json:"w"`
									H int `json:"h"`
								} `json:"focus_rects"`
							} `json:"original_info"`
						} `json:"media"`
					} `json:"extended_entities"`
					FavoriteCount             int    `json:"favorite_count"`
					Favorited                 bool   `json:"favorited"`
					FullText                  string `json:"full_text"`
					IsQuoteStatus             bool   `json:"is_quote_status"`
					Lang                      string `json:"lang"`
					PossiblySensitive         bool   `json:"possibly_sensitive"`
					PossiblySensitiveEditable bool   `json:"possibly_sensitive_editable"`
					QuoteCount                int    `json:"quote_count"`
					ReplyCount                int    `json:"reply_count"`
					RetweetCount              int    `json:"retweet_count"`
					Retweeted                 bool   `json:"retweeted"`
					UserIDStr                 string `json:"user_id_str"`
					IDStr                     string `json:"id_str"`
				} `json:"legacy"`
			} `json:"result"`
		} `json:"tweetResult"`
	} `json:"data"`
}
