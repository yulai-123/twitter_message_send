package twitter

type PostListResp struct {
	Data struct {
		User struct {
			Result struct {
				Typename   string `json:"__typename"`
				TimelineV2 struct {
					Timeline struct {
						Instructions []struct {
							Type  string `json:"type"`
							Entry struct {
								EntryID   string `json:"entryId"`
								SortIndex string `json:"sortIndex"`
								Content   struct {
									EntryType   string `json:"entryType"`
									Typename    string `json:"__typename"`
									ItemContent struct {
										ItemType     string `json:"itemType"`
										Typename     string `json:"__typename"`
										TweetResults struct {
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
															HasGraduatedAccess bool   `json:"has_graduated_access"`
															IsBlueVerified     bool   `json:"is_blue_verified"`
															ProfileImageShape  string `json:"profile_image_shape"`
															Legacy             struct {
																Following           bool   `json:"following"`
																CanDm               bool   `json:"can_dm"`
																CanMediaTag         bool   `json:"can_media_tag"`
																CreatedAt           string `json:"created_at"`
																DefaultProfile      bool   `json:"default_profile"`
																DefaultProfileImage bool   `json:"default_profile_image"`
																Description         string `json:"description"`
																Entities            struct {
																	Description struct {
																		Urls []struct {
																			DisplayURL  string `json:"display_url"`
																			ExpandedURL string `json:"expanded_url"`
																			URL         string `json:"url"`
																			Indices     []int  `json:"indices"`
																		} `json:"urls"`
																	} `json:"description"`
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
																PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
																PossiblySensitive       bool          `json:"possibly_sensitive"`
																ProfileBannerURL        string        `json:"profile_banner_url"`
																ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
																ProfileInterstitialType string        `json:"profile_interstitial_type"`
																ScreenName              string        `json:"screen_name"`
																StatusesCount           int           `json:"statuses_count"`
																TranslatorType          string        `json:"translator_type"`
																Verified                bool          `json:"verified"`
																WantRetweets            bool          `json:"want_retweets"`
																WithheldInCountries     []interface{} `json:"withheld_in_countries"`
															} `json:"legacy"`
														} `json:"result"`
													} `json:"user_results"`
												} `json:"core"`
												UnmentionData struct {
												} `json:"unmention_data"`
												EditControl struct {
													InitialTweetID     string `json:"initial_tweet_id"`
													EditControlInitial struct {
														EditTweetIds       []string `json:"edit_tweet_ids"`
														EditableUntilMsecs string   `json:"editable_until_msecs"`
														IsEditEligible     bool     `json:"is_edit_eligible"`
														EditsRemaining     string   `json:"edits_remaining"`
													} `json:"edit_control_initial"`
												} `json:"edit_control"`
												PreviousCounts struct {
													BookmarkCount int `json:"bookmark_count"`
													FavoriteCount int `json:"favorite_count"`
													QuoteCount    int `json:"quote_count"`
													ReplyCount    int `json:"reply_count"`
													RetweetCount  int `json:"retweet_count"`
												} `json:"previous_counts"`
												IsTranslatable bool `json:"is_translatable"`
												Views          struct {
													Count string `json:"count"`
													State string `json:"state"`
												} `json:"views"`
												Source    string `json:"source"`
												NoteTweet struct {
													IsExpandable     bool `json:"is_expandable"`
													NoteTweetResults struct {
														Result struct {
															ID        string `json:"id"`
															Text      string `json:"text"`
															EntitySet struct {
																Hashtags []interface{} `json:"hashtags"`
																Symbols  []interface{} `json:"symbols"`
																Urls     []struct {
																	DisplayURL  string `json:"display_url"`
																	ExpandedURL string `json:"expanded_url"`
																	URL         string `json:"url"`
																	Indices     []int  `json:"indices"`
																} `json:"urls"`
																UserMentions []struct {
																	IDStr      string `json:"id_str"`
																	Name       string `json:"name"`
																	ScreenName string `json:"screen_name"`
																	Indices    []int  `json:"indices"`
																} `json:"user_mentions"`
															} `json:"entity_set"`
															Richtext struct {
																RichtextTags []struct {
																	FromIndex     int      `json:"from_index"`
																	ToIndex       int      `json:"to_index"`
																	RichtextTypes []string `json:"richtext_types"`
																} `json:"richtext_tags"`
															} `json:"richtext"`
															Media struct {
																InlineMedia []interface{} `json:"inline_media"`
															} `json:"media"`
														} `json:"result"`
													} `json:"note_tweet_results"`
												} `json:"note_tweet"`
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
												QuickPromoteEligibility struct {
													Eligibility string `json:"eligibility"`
												} `json:"quick_promote_eligibility"`
											} `json:"result"`
										} `json:"tweet_results"`
										TweetDisplayType string `json:"tweetDisplayType"`
										SocialContext    struct {
											Type        string `json:"type"`
											ContextType string `json:"contextType"`
											Text        string `json:"text"`
										} `json:"socialContext"`
									} `json:"itemContent"`
									ClientEventInfo struct {
										Component string `json:"component"`
										Element   string `json:"element"`
									} `json:"clientEventInfo"`
								} `json:"content"`
							} `json:"entry,omitempty"`
							Entries []struct {
								EntryID   string `json:"entryId"`
								SortIndex string `json:"sortIndex"`
								Content   struct {
									EntryType string `json:"entryType"`
									Typename  string `json:"__typename"`
									Items     []struct {
										EntryID string `json:"entryId"`
										Item    struct {
											ItemContent struct {
												ItemType     string `json:"itemType"`
												Typename     string `json:"__typename"`
												TweetResults struct {
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
																	HasGraduatedAccess bool   `json:"has_graduated_access"`
																	IsBlueVerified     bool   `json:"is_blue_verified"`
																	ProfileImageShape  string `json:"profile_image_shape"`
																	Legacy             struct {
																		Following           bool   `json:"following"`
																		CanDm               bool   `json:"can_dm"`
																		CanMediaTag         bool   `json:"can_media_tag"`
																		CreatedAt           string `json:"created_at"`
																		DefaultProfile      bool   `json:"default_profile"`
																		DefaultProfileImage bool   `json:"default_profile_image"`
																		Description         string `json:"description"`
																		Entities            struct {
																			Description struct {
																				Urls []struct {
																					DisplayURL  string `json:"display_url"`
																					ExpandedURL string `json:"expanded_url"`
																					URL         string `json:"url"`
																					Indices     []int  `json:"indices"`
																				} `json:"urls"`
																			} `json:"description"`
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
																		PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
																		PossiblySensitive       bool          `json:"possibly_sensitive"`
																		ProfileBannerURL        string        `json:"profile_banner_url"`
																		ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
																		ProfileInterstitialType string        `json:"profile_interstitial_type"`
																		ScreenName              string        `json:"screen_name"`
																		StatusesCount           int           `json:"statuses_count"`
																		TranslatorType          string        `json:"translator_type"`
																		Verified                bool          `json:"verified"`
																		WantRetweets            bool          `json:"want_retweets"`
																		WithheldInCountries     []interface{} `json:"withheld_in_countries"`
																	} `json:"legacy"`
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
															Count string `json:"count"`
															State string `json:"state"`
														} `json:"views"`
														Source    string `json:"source"`
														NoteTweet struct {
															IsExpandable     bool `json:"is_expandable"`
															NoteTweetResults struct {
																Result struct {
																	ID        string `json:"id"`
																	Text      string `json:"text"`
																	EntitySet struct {
																		Hashtags []interface{} `json:"hashtags"`
																		Symbols  []interface{} `json:"symbols"`
																		Urls     []struct {
																			DisplayURL  string `json:"display_url"`
																			ExpandedURL string `json:"expanded_url"`
																			URL         string `json:"url"`
																			Indices     []int  `json:"indices"`
																		} `json:"urls"`
																		UserMentions []struct {
																			IDStr      string `json:"id_str"`
																			Name       string `json:"name"`
																			ScreenName string `json:"screen_name"`
																			Indices    []int  `json:"indices"`
																		} `json:"user_mentions"`
																	} `json:"entity_set"`
																	Richtext struct {
																		RichtextTags []interface{} `json:"richtext_tags"`
																	} `json:"richtext"`
																	Media struct {
																		InlineMedia []interface{} `json:"inline_media"`
																	} `json:"media"`
																} `json:"result"`
															} `json:"note_tweet_results"`
														} `json:"note_tweet"`
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
																			Faces []struct {
																				X int `json:"x"`
																				Y int `json:"y"`
																				H int `json:"h"`
																				W int `json:"w"`
																			} `json:"faces"`
																		} `json:"large"`
																		Medium struct {
																			Faces []struct {
																				X int `json:"x"`
																				Y int `json:"y"`
																				H int `json:"h"`
																				W int `json:"w"`
																			} `json:"faces"`
																		} `json:"medium"`
																		Small struct {
																			Faces []struct {
																				X int `json:"x"`
																				Y int `json:"y"`
																				H int `json:"h"`
																				W int `json:"w"`
																			} `json:"faces"`
																		} `json:"small"`
																		Orig struct {
																			Faces []struct {
																				X int `json:"x"`
																				Y int `json:"y"`
																				H int `json:"h"`
																				W int `json:"w"`
																			} `json:"faces"`
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
																			Faces []struct {
																				X int `json:"x"`
																				Y int `json:"y"`
																				H int `json:"h"`
																				W int `json:"w"`
																			} `json:"faces"`
																		} `json:"large"`
																		Medium struct {
																			Faces []struct {
																				X int `json:"x"`
																				Y int `json:"y"`
																				H int `json:"h"`
																				W int `json:"w"`
																			} `json:"faces"`
																		} `json:"medium"`
																		Small struct {
																			Faces []struct {
																				X int `json:"x"`
																				Y int `json:"y"`
																				H int `json:"h"`
																				W int `json:"w"`
																			} `json:"faces"`
																		} `json:"small"`
																		Orig struct {
																			Faces []struct {
																				X int `json:"x"`
																				Y int `json:"y"`
																				H int `json:"h"`
																				W int `json:"w"`
																			} `json:"faces"`
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
														QuickPromoteEligibility struct {
															Eligibility string `json:"eligibility"`
														} `json:"quick_promote_eligibility"`
													} `json:"result"`
												} `json:"tweet_results"`
												TweetDisplayType string `json:"tweetDisplayType"`
											} `json:"itemContent"`
											ClientEventInfo struct {
												Component string `json:"component"`
												Element   string `json:"element"`
												Details   struct {
													TimelinesDetails struct {
														InjectionType  string `json:"injectionType"`
														ControllerData string `json:"controllerData"`
													} `json:"timelinesDetails"`
												} `json:"details"`
											} `json:"clientEventInfo"`
										} `json:"item"`
									} `json:"items"`
									ItemContent struct {
										ItemType     string `json:"itemType"`
										Typename     string `json:"__typename"`
										TweetResults struct {
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
															HasGraduatedAccess bool   `json:"has_graduated_access"`
															IsBlueVerified     bool   `json:"is_blue_verified"`
															ProfileImageShape  string `json:"profile_image_shape"`
															Legacy             struct {
																Following           bool   `json:"following"`
																CanDm               bool   `json:"can_dm"`
																CanMediaTag         bool   `json:"can_media_tag"`
																CreatedAt           string `json:"created_at"`
																DefaultProfile      bool   `json:"default_profile"`
																DefaultProfileImage bool   `json:"default_profile_image"`
																Description         string `json:"description"`
																Entities            struct {
																	Description struct {
																		Urls []struct {
																			DisplayURL  string `json:"display_url"`
																			ExpandedURL string `json:"expanded_url"`
																			URL         string `json:"url"`
																			Indices     []int  `json:"indices"`
																		} `json:"urls"`
																	} `json:"description"`
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
																PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
																PossiblySensitive       bool          `json:"possibly_sensitive"`
																ProfileBannerURL        string        `json:"profile_banner_url"`
																ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
																ProfileInterstitialType string        `json:"profile_interstitial_type"`
																ScreenName              string        `json:"screen_name"`
																StatusesCount           int           `json:"statuses_count"`
																TranslatorType          string        `json:"translator_type"`
																Verified                bool          `json:"verified"`
																WantRetweets            bool          `json:"want_retweets"`
																WithheldInCountries     []interface{} `json:"withheld_in_countries"`
															} `json:"legacy"`
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
													Count string `json:"count"`
													State string `json:"state"`
												} `json:"views"`
												Source             string `json:"source"`
												QuotedStatusResult struct {
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
																	HasGraduatedAccess bool   `json:"has_graduated_access"`
																	IsBlueVerified     bool   `json:"is_blue_verified"`
																	ProfileImageShape  string `json:"profile_image_shape"`
																	Legacy             struct {
																		CanDm               bool   `json:"can_dm"`
																		CanMediaTag         bool   `json:"can_media_tag"`
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
																		PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
																		PossiblySensitive       bool          `json:"possibly_sensitive"`
																		ProfileBannerURL        string        `json:"profile_banner_url"`
																		ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
																		ProfileInterstitialType string        `json:"profile_interstitial_type"`
																		ScreenName              string        `json:"screen_name"`
																		StatusesCount           int           `json:"statuses_count"`
																		TranslatorType          string        `json:"translator_type"`
																		URL                     string        `json:"url"`
																		Verified                bool          `json:"verified"`
																		WantRetweets            bool          `json:"want_retweets"`
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
															Count string `json:"count"`
															State string `json:"state"`
														} `json:"views"`
														Source    string `json:"source"`
														NoteTweet struct {
															IsExpandable     bool `json:"is_expandable"`
															NoteTweetResults struct {
																Result struct {
																	ID        string `json:"id"`
																	Text      string `json:"text"`
																	EntitySet struct {
																		Hashtags     []interface{} `json:"hashtags"`
																		Symbols      []interface{} `json:"symbols"`
																		Urls         []interface{} `json:"urls"`
																		UserMentions []struct {
																			IDStr      string `json:"id_str"`
																			Name       string `json:"name"`
																			ScreenName string `json:"screen_name"`
																			Indices    []int  `json:"indices"`
																		} `json:"user_mentions"`
																	} `json:"entity_set"`
																} `json:"result"`
															} `json:"note_tweet_results"`
														} `json:"note_tweet"`
														QuotedRefResult struct {
															Result struct {
																Typename string `json:"__typename"`
																RestID   string `json:"rest_id"`
															} `json:"result"`
														} `json:"quotedRefResult"`
														Legacy struct {
															BookmarkCount     int    `json:"bookmark_count"`
															Bookmarked        bool   `json:"bookmarked"`
															CreatedAt         string `json:"created_at"`
															ConversationIDStr string `json:"conversation_id_str"`
															DisplayTextRange  []int  `json:"display_text_range"`
															Entities          struct {
																Hashtags     []interface{} `json:"hashtags"`
																Symbols      []interface{} `json:"symbols"`
																Timestamps   []interface{} `json:"timestamps"`
																Urls         []interface{} `json:"urls"`
																UserMentions []interface{} `json:"user_mentions"`
															} `json:"entities"`
															FavoriteCount         int    `json:"favorite_count"`
															Favorited             bool   `json:"favorited"`
															FullText              string `json:"full_text"`
															IsQuoteStatus         bool   `json:"is_quote_status"`
															Lang                  string `json:"lang"`
															QuoteCount            int    `json:"quote_count"`
															QuotedStatusIDStr     string `json:"quoted_status_id_str"`
															QuotedStatusPermalink struct {
																URL      string `json:"url"`
																Expanded string `json:"expanded"`
																Display  string `json:"display"`
															} `json:"quoted_status_permalink"`
															ReplyCount   int    `json:"reply_count"`
															RetweetCount int    `json:"retweet_count"`
															Retweeted    bool   `json:"retweeted"`
															UserIDStr    string `json:"user_id_str"`
															IDStr        string `json:"id_str"`
														} `json:"legacy"`
													} `json:"result"`
												} `json:"quoted_status_result"`
												Legacy struct {
													BookmarkCount     int    `json:"bookmark_count"`
													Bookmarked        bool   `json:"bookmarked"`
													CreatedAt         string `json:"created_at"`
													ConversationIDStr string `json:"conversation_id_str"`
													DisplayTextRange  []int  `json:"display_text_range"`
													Entities          struct {
														Hashtags     []interface{} `json:"hashtags"`
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
													FavoriteCount         int    `json:"favorite_count"`
													Favorited             bool   `json:"favorited"`
													FullText              string `json:"full_text"`
													IsQuoteStatus         bool   `json:"is_quote_status"`
													Lang                  string `json:"lang"`
													QuoteCount            int    `json:"quote_count"`
													QuotedStatusIDStr     string `json:"quoted_status_id_str"`
													QuotedStatusPermalink struct {
														URL      string `json:"url"`
														Expanded string `json:"expanded"`
														Display  string `json:"display"`
													} `json:"quoted_status_permalink"`
													ReplyCount   int    `json:"reply_count"`
													RetweetCount int    `json:"retweet_count"`
													Retweeted    bool   `json:"retweeted"`
													UserIDStr    string `json:"user_id_str"`
													IDStr        string `json:"id_str"`
												} `json:"legacy"`
												QuickPromoteEligibility struct {
													Eligibility string `json:"eligibility"`
												} `json:"quick_promote_eligibility"`
											} `json:"result"`
										} `json:"tweet_results"`
										TweetDisplayType string `json:"tweetDisplayType"`
									} `json:"itemContent"`
									ClientEventInfo struct {
										Component string `json:"component"`
										Element   string `json:"element"`
										Details   struct {
											TimelinesDetails struct {
												InjectionType  string `json:"injectionType"`
												ControllerData string `json:"controllerData"`
											} `json:"timelinesDetails"`
										} `json:"details"`
									} `json:"clientEventInfo"`
								} `json:"content,omitempty"`
							} `json:"entries,omitempty"`
						} `json:"instructions"`
						Metadata struct {
							ScribeConfig struct {
								Page string `json:"page"`
							} `json:"scribeConfig"`
						} `json:"metadata"`
					} `json:"timeline"`
				} `json:"timeline_v2"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}
