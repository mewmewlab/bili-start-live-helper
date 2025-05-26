package main

const (
	HEADER_USER_AGENT = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"
)

const (
	QRCODE_URL       = "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	CHECK_QRCODE_URL = "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"

	USER_INFO_URL = "https://api.bilibili.com/x/space/v2/myinfo"
	GO_URL        = "https://live.bilibili.com/"
	ORIGIN_URL    = "https://passport.bilibili.com"
	ROOMINFO_URL  = "https://api.live.bilibili.com/xlive/web-ucenter/user/live_info"

	SEND_DANMAKU_URL     = "https://api.live.bilibili.com/msg/send"
	LIVE_CATEGORY_URL    = "https://api.live.bilibili.com/room/v1/Area/getList?show_pinyin=1"
	UPDATE_TITLE_URL     = "https://api.live.bilibili.com/room/v1/Room/update"
	START_LIVE_URL       = "https://api.live.bilibili.com/room/v1/Room/startLive"
	STOP_LIVE_URL        = "https://api.live.bilibili.com/room/v1/Room/stopLive"
	LIVE_ROOM_STATUS_URL = "https://api.live.bilibili.com/room/v1/Room/get_info"

	COOKIES_SAVE_PATH = "cookies"
)

const (
	TimeoutOrNotMatch = -2
	NotScan           = -4
	NotConfirm        = -5
)

const (
	QRCODE_LOGIN_SUCCESS = 0
	QRCODE_TIMEOUT       = 86038
	QRCODE_WAIT_CONFIRM  = 86090
	QRCODE_NOT_SCAN      = 86101
)

const (
	NOT_LOGIN = -101
)
