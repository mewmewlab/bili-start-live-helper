package model

type UserInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Profile struct {
			Mid            int    `json:"mid"`
			Name           string `json:"name"`
			Sex            string `json:"sex"`
			Face           string `json:"face"`
			Sign           string `json:"sign"`
			Rank           int    `json:"rank"`
			Level          int    `json:"level"`
			Jointime       int    `json:"jointime"`
			Moral          int    `json:"moral"`
			Silence        int    `json:"silence"`
			EmailStatus    int    `json:"email_status"`
			TelStatus      int    `json:"tel_status"`
			Identification int    `json:"identification"`
			Vip            struct {
				Type       int   `json:"type"`
				Status     int   `json:"status"`
				DueDate    int64 `json:"due_date"`
				VipPayType int   `json:"vip_pay_type"`
				ThemeType  int   `json:"theme_type"`
				Label      struct {
					Path                  string `json:"path"`
					Text                  string `json:"text"`
					LabelTheme            string `json:"label_theme"`
					TextColor             string `json:"text_color"`
					BgStyle               int    `json:"bg_style"`
					BgColor               string `json:"bg_color"`
					BorderColor           string `json:"border_color"`
					UseImgLabel           bool   `json:"use_img_label"`
					ImgLabelURIHans       string `json:"img_label_uri_hans"`
					ImgLabelURIHant       string `json:"img_label_uri_hant"`
					ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
					ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				Role               int    `json:"role"`
				AvatarSubscriptURL string `json:"avatar_subscript_url"`
				TvVipStatus        int    `json:"tv_vip_status"`
				TvVipPayType       int    `json:"tv_vip_pay_type"`
				TvDueDate          int    `json:"tv_due_date"`
				AvatarIcon         struct {
					IconType     int `json:"icon_type"`
					IconResource struct {
					} `json:"icon_resource"`
				} `json:"avatar_icon"`
			} `json:"vip"`
			Pendant struct {
				Pid               int    `json:"pid"`
				Name              string `json:"name"`
				Image             string `json:"image"`
				Expire            int    `json:"expire"`
				ImageEnhance      string `json:"image_enhance"`
				ImageEnhanceFrame string `json:"image_enhance_frame"`
				NPid              int    `json:"n_pid"`
			} `json:"pendant"`
			Nameplate struct {
				Nid        int    `json:"nid"`
				Name       string `json:"name"`
				Image      string `json:"image"`
				ImageSmall string `json:"image_small"`
				Level      string `json:"level"`
				Condition  string `json:"condition"`
			} `json:"nameplate"`
			Official struct {
				Role  int    `json:"role"`
				Title string `json:"title"`
				Desc  string `json:"desc"`
				Type  int    `json:"type"`
			} `json:"official"`
			Birthday      int  `json:"birthday"`
			IsTourist     int  `json:"is_tourist"`
			IsFakeAccount int  `json:"is_fake_account"`
			PinPrompting  int  `json:"pin_prompting"`
			IsDeleted     int  `json:"is_deleted"`
			InRegAudit    int  `json:"in_reg_audit"`
			IsRipUser     bool `json:"is_rip_user"`
			Profession    struct {
				ID              int    `json:"id"`
				Name            string `json:"name"`
				ShowName        string `json:"show_name"`
				IsShow          int    `json:"is_show"`
				CategoryOne     string `json:"category_one"`
				Realname        string `json:"realname"`
				Title           string `json:"title"`
				Department      string `json:"department"`
				CertificateNo   string `json:"certificate_no"`
				CertificateShow bool   `json:"certificate_show"`
			} `json:"profession"`
			FaceNft        int `json:"face_nft"`
			FaceNftNew     int `json:"face_nft_new"`
			IsSeniorMember int `json:"is_senior_member"`
			Honours        struct {
				Mid    int `json:"mid"`
				Colour struct {
					Dark   string `json:"dark"`
					Normal string `json:"normal"`
				} `json:"colour"`
				Tags              any `json:"tags"`
				IsLatest100Honour int `json:"is_latest_100honour"`
			} `json:"honours"`
			DigitalID   string `json:"digital_id"`
			DigitalType int    `json:"digital_type"`
			Attestation struct {
				Type       int `json:"type"`
				CommonInfo struct {
					Title       string `json:"title"`
					Prefix      string `json:"prefix"`
					PrefixTitle string `json:"prefix_title"`
				} `json:"common_info"`
				SpliceInfo struct {
					Title string `json:"title"`
				} `json:"splice_info"`
				Icon string `json:"icon"`
				Desc string `json:"desc"`
			} `json:"attestation"`
			ExpertInfo struct {
				Title string `json:"title"`
				State int    `json:"state"`
				Type  int    `json:"type"`
				Desc  string `json:"desc"`
			} `json:"expert_info"`
			NameRender  any    `json:"name_render"`
			CountryCode string `json:"country_code"`
		} `json:"profile"`
		LevelExp struct {
			CurrentLevel int   `json:"current_level"`
			CurrentMin   int   `json:"current_min"`
			CurrentExp   int   `json:"current_exp"`
			NextExp      int   `json:"next_exp"`
			LevelUp      int64 `json:"level_up"`
		} `json:"level_exp"`
		Coins     float64 `json:"coins"`
		Following int     `json:"following"`
		Follower  int     `json:"follower"`
	} `json:"data,omitempty"`
	TTL int `json:"ttl"`
}

type RoomInfo struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	TTL     int    `json:"ttl,omitempty"`
	Data    struct {
		RoomID            int `json:"room_id,omitempty"`
		MainSiteLevelRank int `json:"main_site_level_rank,omitempty"`
		Master            struct {
			Level   int `json:"level,omitempty"`
			Current int `json:"current,omitempty"`
			Next    int `json:"next,omitempty"`
			Medal   struct {
				ID             int    `json:"id,omitempty"`
				UID            int    `json:"uid,omitempty"`
				MedalName      string `json:"medal_name,omitempty"`
				LiveStatus     int    `json:"live_status,omitempty"`
				MasterStatus   int    `json:"master_status,omitempty"`
				Status         int    `json:"status,omitempty"`
				Reason         string `json:"reason,omitempty"`
				LastWearTime   int    `json:"last_wear_time,omitempty"`
				RenameStatus   int    `json:"rename_status,omitempty"`
				TimeAbleChange int    `json:"time_able_change,omitempty"`
				RenameReason   string `json:"rename_reason,omitempty"`
				ApplySource    int    `json:"apply_source,omitempty"`
				ChargeNum      int    `json:"charge_num,omitempty"`
				CoinNum        int    `json:"coin_num,omitempty"`
				PlatformStatus int    `json:"platform_status,omitempty"`
			} `json:"medal,omitempty"`
		} `json:"master,omitempty"`
		VipInfo struct {
			VipEndtime    string `json:"vip_endtime,omitempty"`
			SvipEndtime   string `json:"svip_endtime,omitempty"`
			MonthPrice    int    `json:"month_price,omitempty"`
			YearPrice     int    `json:"year_price,omitempty"`
			VipViewStatus bool   `json:"vip_view_status,omitempty"`
		} `json:"vip_info,omitempty"`
		LiveTime  int `json:"live_time,omitempty"`
		BiliCoins int `json:"bili_coins,omitempty"`
		San       int `json:"san,omitempty"`
		Count     struct {
			Guard     int `json:"guard,omitempty"`
			FansMedal int `json:"fans_medal,omitempty"`
			Title     int `json:"title,omitempty"`
			TitleNew  int `json:"title_new,omitempty"`
			Achieve   int `json:"achieve,omitempty"`
		} `json:"count,omitempty"`
		AnchorSwitchInfo struct {
			VoiceBarrage int `json:"voice_barrage,omitempty"`
		} `json:"anchor_switch_info,omitempty"`
		NeedGuide              int `json:"need_guide,omitempty"`
		EmoticonManagement     int `json:"emoticon_management,omitempty"`
		UpEmoticonJurisdiction int `json:"up_emoticon_jurisdiction,omitempty"`
		AnchorRoomEmoticon     int `json:"anchor_room_emoticon,omitempty"`
		AnchorViolationsRecord int `json:"anchor_violations_record,omitempty"`
	} `json:"data,omitempty"`
}
