package model

// QRCode 二维码
type QRCode struct {
	Code    int    `json:"code,omitempty"`
	TTL     int    `json:"ttl,omitempty"`
	Message string `json:"message,omitempty"`
	Data    Data   `json:"data"`
}

// Data 二维码数据
type Data struct {
	URL       string `json:"url"`
	QRCodeKey string `json:"qrcode_key"`
}

// CheckLogin 检查登陆
type CheckLogin struct {
	Code    int            `json:"code,omitempty"`
	TTL     int            `json:"ttl,omitempty"`
	Message string         `json:"message,omitempty"`
	Data    CheckLoginData `json:"data"`
}

type CheckLoginData struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	RefreshToken string `json:"refresh_token"`
	Timestamp    int    `json:"timestamp"`
	URL          string `json:"url"`
}
