package main

import (
	"bili-start-live-helper/model"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	cookiejar "github.com/juju/persistent-cookiejar"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/net/publicsuffix"
	"resty.dev/v3"
)

// App struct
type App struct {
	ctx context.Context

	jar     *cookiejar.Jar
	client  *http.Client
	client2 *resty.Client

	csrf       string
	session_id string
	userInfo   *model.UserInfo
	roomInfo   *model.RoomInfo
}

// NewApp creates a new App application struct
func NewApp() *App {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(homedir, ".bili-start-live-helper")
	os.MkdirAll(configPath, os.ModePerm)
	jar, _ := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
		Filename:         filepath.Join(configPath, COOKIES_SAVE_PATH),
	})
	client := &http.Client{
		Jar: jar,
	}
	a := &App{
		jar:    jar,
		client: client,
		client2: resty.New().
			SetHeader("User-Agent", HEADER_USER_AGENT).
			SetBaseURL(GO_URL).SetCookieJar(jar),
	}
	return a
}

func (a *App) GetQRCode() (model.QRCode, error) {
	var qr model.QRCode
	resp, err := a.client2.R().Get(QRCODE_URL)
	if err != nil {
		return qr, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &qr); err != nil {
		return qr, err
	}
	go a.getLoginStatus(a.ctx, qr.Data.QRCodeKey)
	return qr, nil
}

func (a *App) saveCookies() error {
	if err := a.jar.Save(); err != nil {
		return err
	}
	a.loadCSRF()
	return nil
}

func (a *App) loadCSRF() {
	ck := a.jar.AllCookies()
	for _, cookie := range ck {
		runtime.LogDebugf(a.ctx, "Cookie: %s = %s", cookie.Name, cookie.Value)
		if cookie.Name == "bili_jct" {
			a.csrf = cookie.Value
		} else if cookie.Name == "SESSDATA" {
			a.session_id = cookie.Value
		}
	}
}

func (app *App) checkLogin(oauth string) (model.CheckLogin, error) {
	var check model.CheckLogin
	resp, err := app.client2.R().SetQueryParams(map[string]string{
		"qrcode_key": oauth}).Get(CHECK_QRCODE_URL)
	if err != nil {
		return check, err
	}
	defer resp.Body.Close()
	block, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(block, &check)
	return check, err
}

func (a *App) getLoginStatus(ctx context.Context, oauth string) {
	for {
		eventKey := oauth + "-status"
		check, err := a.checkLogin(oauth)
		if err != nil {
			runtime.LogError(ctx, err.Error())
			goto next
		}
		switch check.Data.Code {
		case QRCODE_LOGIN_SUCCESS:
			runtime.LogInfof(ctx, "Login success")
			if err := a.saveCookies(); err != nil {
				runtime.LogErrorf(ctx, "Save cookies error: %v, Please re-login", err.Error())
				runtime.EventsEmit(ctx, eventKey, -1) // -1 means relogin
			} else {
				a.loadUserInfo()
				a.loadRoomInfo()
				runtime.EventsEmit(ctx, eventKey, 3) // 3 means success
			}
			goto exit
		case QRCODE_TIMEOUT:
			runtime.LogWarning(ctx, "Qrcode timeout")
			runtime.EventsEmit(ctx, eventKey, 1) // 1 means timeout
			goto exit
		case QRCODE_WAIT_CONFIRM:
			runtime.LogInfo(ctx, "Waiting for tel confirm")
			runtime.EventsEmit(ctx, eventKey, 2) // 2 means waitting tel confirm
		default:
			runtime.EventsEmit(ctx, eventKey, 0) // 0 means waiting scan
		}
	next:
		time.Sleep(time.Second * time.Duration(2))
	}
exit:
}

func (a *App) loadUserInfo() {
	info := &model.UserInfo{}
	resp, err := a.client2.R().Get(USER_INFO_URL)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &info); err != nil {
		runtime.LogError(a.ctx, err.Error())
		return
	}
	if info.Code == NOT_LOGIN {
		return
	}
	a.userInfo = info
}

func (a *App) loadRoomInfo() {
	var room model.RoomInfo
	resp, err := a.client2.R().Get(ROOMINFO_URL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &room); err != nil {
		runtime.LogError(a.ctx, err.Error())
		return
	}
	a.roomInfo = &room
}

type User struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (a *App) UserInfo() (*User, error) {
	if a.userInfo == nil || a.roomInfo == nil {
		return nil, errors.New("not login")
	}
	return &User{
		Name:   a.userInfo.Data.Profile.Name,
		Avatar: a.userInfo.Data.Profile.Face,
		RoomID: a.roomInfo.Data.RoomID,
	}, nil
}

// ============= 获取直播间分类数据 ==============
type LiveCategory struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    []CategoryItem `json:"data"`
}

type CategoryItem struct {
	ID   int               `json:"id"`
	Name string            `json:"name"`
	List []CategorySubItem `json:"list"`
}

type CategorySubItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (a *App) GetLiveCategory() LiveCategory {
	var roomCategory model.LiveCategory
	resp, err := a.client2.R().Get(LIVE_CATEGORY_URL)
	if err != nil {
		return LiveCategory{
			Code:    -1,
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &roomCategory); err != nil {
		return LiveCategory{
			Code:    -1,
			Message: err.Error(),
		}
	}
	var res LiveCategory
	res.Data = make([]CategoryItem, len(roomCategory.Data))
	for idx, cat := range roomCategory.Data {
		subcat := CategoryItem{
			ID:   cat.ID,
			Name: cat.Name,
			List: make([]CategorySubItem, len(cat.List)),
		}
		for subidx, subsubcat := range cat.List {
			intID, _ := strconv.Atoi(subsubcat.ID)
			subcat.List[subidx] = CategorySubItem{
				ID:   intID,
				Name: subsubcat.Name,
			}
		}
		res.Data[idx] = subcat
	}
	return res
}

// ============= 开始直播==============
func (a *App) StartLive(title string, catID int) (model.Rtmp, error) {
	runtime.LogDebugf(a.ctx, "Start live: %s, catID: %d", title, catID)
	runtime.LogDebugf(a.ctx, "Start Live Params: CSRF = %s, RoomID = %d, Title = %s, catID = %d", a.csrf, a.roomInfo.Data.RoomID, title, catID)
	runtime.LogDebugf(a.ctx, "Cookies: %v", a.jar.AllCookies())
	var startLive model.StartLive
	resp, err := a.client2.R().SetQueryParams(map[string]string{
		"room_id":  strconv.Itoa(a.roomInfo.Data.RoomID),
		"csrf":     a.csrf,
		"platform": "pc_link",
		"area_v2":  strconv.Itoa(catID),
	}).Post(START_LIVE_URL)
	if err != nil {
		return startLive.Data.Rtmp, err
	}
	defer resp.Body.Close()
	jdata, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(jdata, &startLive); err != nil {
		return startLive.Data.Rtmp, err
	}
	if startLive.Code != 0 {
		runtime.LogErrorf(a.ctx, "Start live error: %v", startLive)
		return startLive.Data.Rtmp, errors.New(startLive.Message)
	}
	return startLive.Data.Rtmp, nil
}

// ============= 更新直播信息 ==============
func (a *App) UpdateLiveInfo(title string, catID int) error {
	var updateInfo model.UpdateRoomInfo
	resp, err := a.client2.R().SetQueryParams(map[string]string{
		"csrf":    a.csrf,
		"room_id": strconv.Itoa(a.roomInfo.Data.RoomID),
		"title":   title,
		"area_id": strconv.Itoa(catID),
	}).Post(UPDATE_TITLE_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	jdata, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(jdata, &updateInfo); err != nil {
		return err
	}
	if updateInfo.Code != 0 {
		return errors.New(updateInfo.Message)
	}
	return nil
}

// ============= 关闭直播==============
func (a *App) StopLive() error {
	var stopLive model.StopLive
	resp, err := a.client2.R().SetQueryParams(map[string]string{
		"room_id": strconv.Itoa(a.roomInfo.Data.RoomID),
		"csrf":    a.csrf,
	}).Post(STOP_LIVE_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	jdata, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(jdata, &stopLive); err != nil {
		return err
	}
	if stopLive.Code != 0 {
		return errors.New(stopLive.Message)
	}
	return nil
}

// ============= 获取直播间状态 =============
type RoomStatus struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Status   int    `json:"status"`
	AreaID   int    `json:"area_id"`
	ParentID int    `json:"parent_area_id"`
	Title    string `json:"title"`
}

func (a *App) GetRoomStatus(roomID int) (*RoomStatus, error) {
	var liveRoomStatus model.LiveRoomStatus
	resp, err := a.client2.R().SetQueryParams(map[string]string{
		"room_id": strconv.Itoa(roomID),
	}).Get(LIVE_ROOM_STATUS_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(data, &liveRoomStatus); err != nil {
		return nil, err
	}
	return &RoomStatus{
		Code:     liveRoomStatus.Code,
		Message:  liveRoomStatus.Message,
		Status:   liveRoomStatus.Data.LiveStatus,
		AreaID:   liveRoomStatus.Data.AreaID,
		ParentID: liveRoomStatus.Data.ParentAreaID,
		Title:    liveRoomStatus.Data.Title,
	}, nil
}

// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	if _, err := os.Stat(COOKIES_SAVE_PATH); err == nil {
		a.loadCSRF()
		a.loadUserInfo()
		a.loadRoomInfo()
	}
}
