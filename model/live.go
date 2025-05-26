package model

type LiveCategory struct {
	Code    int    `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		List []struct {
			ID              string `json:"id,omitempty"`
			ParentID        string `json:"parent_id,omitempty"`
			OldAreaID       string `json:"old_area_id,omitempty"`
			Name            string `json:"name,omitempty"`
			ActID           string `json:"act_id,omitempty"`
			PkStatus        string `json:"pk_status,omitempty"`
			HotStatus       int    `json:"hot_status,omitempty"`
			LockStatus      string `json:"lock_status,omitempty"`
			Pic             string `json:"pic,omitempty"`
			ComplexAreaName string `json:"complex_area_name,omitempty"`
			Pinyin          string `json:"pinyin,omitempty"`
			ParentName      string `json:"parent_name,omitempty"`
			AreaType        int    `json:"area_type,omitempty"`
			CateID          string `json:"cate_id,omitempty"`
		} `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

type StartLive struct {
	Code    int           `json:"code,omitempty"`
	TTL     int           `json:"ttl,omitempty"`
	Message string        `json:"message,omitempty"`
	Data    StartLiveData `json:"data"`
}

type StartLiveData struct {
	Change        int           `json:"change"`
	Status        string        `json:"status"`
	RoomType      int           `json:"room_type"`
	Rtmp          Rtmp          `json:"rtmp"`
	Protocols     []interface{} `json:"protocols"`
	TryTime       string        `json:"try_time"`
	LiveKey       string        `json:"live_key"`
	SubSessionKey string        `json:"sub_session_key"`
	Notice        Notice        `json:"notice"`
	QR            string        `json:"qr"`
	NeedFaceAuth  bool          `json:"need_face_auth"`
	ServiceSource string        `json:"service_source"`
	RTMPBackup    interface{}   `json:"rtmp_backup"`
	UpStreamExtra interface{}   `json:"up_stream_extra"`
}

type Rtmp struct {
	Addr     string `json:"addr"`
	Code     string `json:"code"`
	NewLink  string `json:"new_link"`
	Provider string `json:"provider"`
}

type Protocols struct {
	Protocol string `json:"protocol"`
	Addr     string `json:"addr"`
	Code     string `json:"code"`
	NewLink  string `json:"new_link"`
	Provider string `json:"provider"`
}

type Notice struct {
	Type       int    `json:"type"`
	Status     int    `json:"status"`
	Title      string `json:"title"`
	Msg        string `json:"msg"`
	ButtonText string `json:"button_text"`
	ButtonURL  string `json:"button_url"`
}

type StopLive struct {
	Code    int          `json:"code,omitempty"`
	TTL     int          `json:"ttl,omitempty"`
	Message string       `json:"message,omitempty"`
	Data    StopLiveData `json:"data"`
}

type StopLiveData struct {
	Change int    `json:"change"`
	Status string `json:"status"`
}

type UpdateRoomInfo struct {
	Code    int                `json:"code,omitempty"`
	TTL     int                `json:"ttl,omitempty"`
	Message string             `json:"message,omitempty"`
	Data    UpdateRoomInfoData `json:"data"`
}

type UpdateRoomInfoData struct {
	SubSessionKey string `json:"sub_session_key"`
	AuditInfo     struct {
		AuditTitleReson  string `json:"audit_title_reson"`
		AuditTitleStatus int    `json:"audit_title_status"`
		AuditTitle       string `json:"audit_title"`
		UpdateTitle      string `json:"update_title"`
	} `json:"audit_info"`
}

type LiveRoomStatus struct {
	Code    int    `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		UID              int      `json:"uid,omitempty"`
		RoomID           int      `json:"room_id,omitempty"`
		ShortID          int      `json:"short_id,omitempty"`
		Attention        int      `json:"attention,omitempty"`
		Online           int      `json:"online,omitempty"`
		IsPortrait       bool     `json:"is_portrait,omitempty"`
		Description      string   `json:"description,omitempty"`
		LiveStatus       int      `json:"live_status,omitempty"`
		AreaID           int      `json:"area_id,omitempty"`
		ParentAreaID     int      `json:"parent_area_id,omitempty"`
		ParentAreaName   string   `json:"parent_area_name,omitempty"`
		OldAreaID        int      `json:"old_area_id,omitempty"`
		Background       string   `json:"background,omitempty"`
		Title            string   `json:"title,omitempty"`
		UserCover        string   `json:"user_cover,omitempty"`
		Keyframe         string   `json:"keyframe,omitempty"`
		IsStrictRoom     bool     `json:"is_strict_room,omitempty"`
		LiveTime         string   `json:"live_time,omitempty"`
		Tags             string   `json:"tags,omitempty"`
		IsAnchor         int      `json:"is_anchor,omitempty"`
		RoomSilentType   string   `json:"room_silent_type,omitempty"`
		RoomSilentLevel  int      `json:"room_silent_level,omitempty"`
		RoomSilentSecond int      `json:"room_silent_second,omitempty"`
		AreaName         string   `json:"area_name,omitempty"`
		Pendants         string   `json:"pendants,omitempty"`
		AreaPendants     string   `json:"area_pendants,omitempty"`
		HotWords         []string `json:"hot_words,omitempty"`
		HotWordsStatus   int      `json:"hot_words_status,omitempty"`
		Verify           string   `json:"verify,omitempty"`
		NewPendants      struct {
			Frame struct {
				Name       string `json:"name,omitempty"`
				Value      string `json:"value,omitempty"`
				Position   int    `json:"position,omitempty"`
				Desc       string `json:"desc,omitempty"`
				Area       int    `json:"area,omitempty"`
				AreaOld    int    `json:"area_old,omitempty"`
				BgColor    string `json:"bg_color,omitempty"`
				BgPic      string `json:"bg_pic,omitempty"`
				UseOldArea bool   `json:"use_old_area,omitempty"`
			} `json:"frame,omitempty"`
			Badge       any `json:"badge,omitempty"`
			MobileFrame struct {
				Name       string `json:"name,omitempty"`
				Value      string `json:"value,omitempty"`
				Position   int    `json:"position,omitempty"`
				Desc       string `json:"desc,omitempty"`
				Area       int    `json:"area,omitempty"`
				AreaOld    int    `json:"area_old,omitempty"`
				BgColor    string `json:"bg_color,omitempty"`
				BgPic      string `json:"bg_pic,omitempty"`
				UseOldArea bool   `json:"use_old_area,omitempty"`
			} `json:"mobile_frame,omitempty"`
			MobileBadge any `json:"mobile_badge,omitempty"`
		} `json:"new_pendants,omitempty"`
		UpSession            string `json:"up_session,omitempty"`
		PkStatus             int    `json:"pk_status,omitempty"`
		PkID                 int    `json:"pk_id,omitempty"`
		BattleID             int    `json:"battle_id,omitempty"`
		AllowChangeAreaTime  int    `json:"allow_change_area_time,omitempty"`
		AllowUploadCoverTime int    `json:"allow_upload_cover_time,omitempty"`
		StudioInfo           struct {
			Status     int   `json:"status,omitempty"`
			MasterList []any `json:"master_list,omitempty"`
		} `json:"studio_info,omitempty"`
	} `json:"data,omitempty"`
}
