// +build websocket

package cqp

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// TODO: 引入连接池提升性能
func connAPIs(urlStr string, requestHeader http.Header) {
	conn, _, err := ws.DefaultDialer.Dial(urlStr, requestHeader)
	if err != nil {
		log.WithField("remoteAddr", urlStr).
			WithError(err).
			Fatal("连接/api失败")
	}
	log.WithField("url", urlStr).Info("成功连接/api")
	apiConn = conn
	// 接收数据
	//for {
	//	var e apiRet
	//	if err := conn.ReadJSON(&e); err != nil {
	//		log.WithError(err).Fatal("API出错")
	//	}
	//	log.Info(e)
	//}
}

type apiRet struct {
	Status  string          `json:"status"`
	RetCode int32           `json:"retcode"`
	Echo    int             `json:"echo"`
	Data    json.RawMessage `json:"data"`
}

var apiConn *ws.Conn
var apiConnMut sync.Mutex

func getAPI() *ws.Conn {
	apiConnMut.Lock()
	return apiConn
}

func retAPI() {
	apiConnMut.Unlock()
}

type apiPayload struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   int         `json:"echo"`
}

type apiParams map[string]interface{}

// callAPI 调用API
func callAPI(args apiPayload, rets interface{}) int32 {
	conn := getAPI()
	defer retAPI()

	echo := rand.Int()
	args.Echo = echo

	err := conn.WriteJSON(args)
	if err != nil {
		log.WithError(err).Error("发送调用请求失败")
	}

	// 检查调用是否成功
	code, err := checkRet(conn, echo, rets)
	if err != nil {
		log.WithError(err).Error("调用失败")
	}

	if code < 0 { // 只有小于0的错误代码才是酷Q本身的
		return code
	}
	return 0
}

// 接收并检查一般的返回值
func checkRet(conn *ws.Conn, echo int, data interface{}) (int32, error) {
	var resp apiRet
	if err := conn.ReadJSON(&resp); err != nil {
		return resp.RetCode, fmt.Errorf("接收API响应失败: %v", err)
	}
	// TODO: 检查echo值一致性
	switch resp.Status {
	case "ok":
		log.WithField("echo", resp.Echo).Debug("调用成功")
	case "async":
		log.Warning("请求异步处理，具体成功或失败将无法获知。")
	case "failed":
		return resp.RetCode, fmt.Errorf("调用API失败: %d", resp.RetCode)
	}
	if data != nil {
		if err := json.Unmarshal(resp.Data, data); err != nil {
			return resp.RetCode, fmt.Errorf("无法解析JSON: %v", err)
		}
	}
	return resp.RetCode, nil
}

// AddLog 增加运行日志
// 	priority 是Log的优先级，请使用cqp预定义好的值。
// 	logType 是日志类型，酷Q日志窗口将将其显示在日志本身的前一列。
// 	reason 是日志内容
func AddLog(p Priority, logType, reason string) int32 {
	ent := log.WithField("logType", logType)
	switch p {
	case Debug:
		ent.Debug(reason)
	case Info, InfoSuccess, InfoRecv, InfoSend:
		ent.Info(reason)
	case Warning:
		ent.Warning(reason)
	case Error:
		ent.Error(reason)
	case Fatal:
		ent.Fatal(reason)
	default:
		ent.Panic("未知优先级日志：", reason)
	}
	return 0
}

// CanSendImage 能否发送图片
func CanSendImage() bool {
	var resp struct {
		Yes bool `json:"yes"`
	}
	callAPI(apiPayload{Action: "can_send_image"}, &resp)
	return resp.Yes
}

// CanSendRecord 能否发送语音
func CanSendRecord() bool {
	var resp struct {
		Yes bool `json:"yes"`
	}
	callAPI(apiPayload{Action: "can_send_record"}, &resp)
	return resp.Yes
}

// DeleteMsg 撤回消息
func DeleteMsg(msgID int64) int32 {
	return callAPI(apiPayload{
		Action: "delete_msg",
		Params: apiParams{"message_id": msgID},
	}, nil)
}

// GetAppDir 取应用目录
// 返回的路径末尾带"\"，一般用法如下，所以不用关心是否有斜线
// 	os.Open(filepath.Join(cqp.GetAppDir(), "data.db"))
func GetAppDir() string {
	return *appDir
}

// GetCookies 获取cookies
//
// 需要严格授权
func GetCookies() string {
	return GetCookiesV2("")
}

// GetCookiesV2 获取cookies
//
// 需要严格授权
func GetCookiesV2(domain string) string {
	var resp struct {
		Cookies string `json:"cookies"`
	}
	callAPI(apiPayload{
		Action: "get_cookies",
		Params: apiParams{"domain": domain},
	}, &resp)
	return resp.Cookies
}

// GetCSRFToken 获取CSRF Token
//
// 需要严格授权
func GetCSRFToken() int32 {
	var resp struct {
		Token int32 `json:"token"`
	}
	callAPI(apiPayload{Action: "get_csrf_token"}, &resp)
	return resp.Token
}

// GetFriendList 获取好友列表
func GetFriendList() (friendList []FriendInfo) {
	var fl []struct {
		QQ    int64  `json:"user_id"`
		Name  string `json:"nickname"`
		Alias string `json:"remark"`
	}
	callAPI(apiPayload{Action: "get_friend_list"}, &fl)

	friendList = make([]FriendInfo, len(fl))
	for i := range fl {
		friendList[i] = FriendInfo(fl[i])
	}
	return
}

func GetGroupInfo(group int64, noCatch bool) GroupDetail {
	var ret struct {
		GroupID        int64  `json:"group_id"`
		GroupName      string `json:"group_name"`
		MemberCount    int32  `json:"member_count"`
		MaxMemberCount int32  `json:"max_member_count"`
	}
	callAPI(apiPayload{
		Action: "get_group_info",
		Params: apiParams{
			"group_id": group,
			"no_cache": noCatch,
		},
	}, &ret)
	return GroupDetail{
		GroupInfo:    GroupInfo{ID: ret.GroupID, Name: ret.GroupName},
		MembersNum:   ret.MemberCount,
		MaxMemberNum: ret.MaxMemberCount,
	}
}

// GetGroupList 获取群列表
func GetGroupList() (groupList []GroupInfo) {
	var ret []struct {
		ID   int64  `json:"group_id"`
		Name string `json:"group_name"`
	}
	callAPI(apiPayload{Action: "get_group_list"}, &ret)
	groupList = make([]GroupInfo, len(ret))
	for i := range ret {
		groupList[i] = GroupInfo(ret[i])
	}
	return
}

// GetGroupMemberInfo 获取群成员信息
func GetGroupMemberInfo(group, qq int64, noCatch bool) (gm GroupMember) {
	var ret struct {
		Group       int64     `json:"group_id"`
		QQ          int64     `json:"user_id"`
		Name        string    `json:"nickname"`
		Card        string    `json:"card"`
		Gender      int32     `json:"sex"`
		Age         int32     `json:"age"`
		Area        string    `json:"area"`
		JoinTime    time.Time `json:"join_time"`
		LastChat    time.Time `json:"last_sent_time"`
		Level       string    `json:"level"`
		Auth        int32     `json:"role"`
		Bad         bool      `json:"unfriendly"`
		Title       string    `json:"title"`
		TitleLife   time.Time `json:"title_expire_time"`
		CanSetTitle bool      `json:"card_changeable"`
	}
	callAPI(apiPayload{
		Action: "get_group_member_info",
		Params: apiParams{
			"group_id": group,
			"user_id":  qq,
			"no_cache": noCatch,
		},
	}, &ret)
	return GroupMember(ret)
}

// GetGroupMemberList 获取群成员列表
func GetGroupMemberList(group int64) (groupMemberList []GroupMember) {
	var ret []struct {
		Group       int64     `json:"group_id"`
		QQ          int64     `json:"user_id"`
		Name        string    `json:"nickname"`
		Card        string    `json:"card"`
		Gender      int32     `json:"sex"`
		Age         int32     `json:"age"`
		Area        string    `json:"area"`
		JoinTime    time.Time `json:"join_time"`
		LastChat    time.Time `json:"last_sent_time"`
		Level       string    `json:"level"`
		Auth        int32     `json:"role"`
		Bad         bool      `json:"unfriendly"`
		Title       string    `json:"title"`
		TitleLife   time.Time `json:"title_expire_time"`
		CanSetTitle bool      `json:"card_changeable"`
	}
	callAPI(apiPayload{
		Action: "get_group_member_info",
		Params: apiParams{"group_id": group},
	}, &ret)

	groupMemberList = make([]GroupMember, len(ret))
	for i := range ret {
		groupMemberList[i] = GroupMember(ret[i])
	}
	return
}

// GetImage 获取图片
// 参数为CQ码内容，返回值为图片的文件路径
func GetImage(file string) string {
	var ret struct {
		File string `json:"file"`
	}
	callAPI(apiPayload{
		Action: "get_image",
		Params: apiParams{"file": file},
	}, &ret)
	return ret.File
}

//GetLoginNick 获取登录号昵称
func GetLoginNick() string {
	var resp struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
	}
	callAPI(apiPayload{Action: "get_login_info"}, &resp)
	return resp.Nickname
}

//GetLoginQQ 获取登陆号QQ
func GetLoginQQ() int64 {
	var resp struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
	}
	callAPI(apiPayload{Action: "get_login_info"}, &resp)
	return resp.UserID
}

// GetRecord 获取语音
// file参数为CQ码内容，format为插件所需格式，返回语音文件路径
func GetRecord(file, format string) string {
	var ret struct {
		File string `json:"file"`
	}
	callAPI(apiPayload{
		Action: "get_record",
		Params: apiParams{
			"file":       file,
			"out_format": format,
			"full_path":  false,
		},
	}, &ret)
	return ret.File
}

// GetRecordV2 获取语音
// 与GetRecord不同之处在于V2返回绝对路径
func GetRecordV2(file, format string) string {
	var ret struct {
		File string `json:"file"`
	}
	callAPI(apiPayload{
		Action: "get_record",
		Params: apiParams{
			"file":       file,
			"out_format": format,
			"full_path":  true,
		},
	}, &ret)
	return ret.File
}

//// GetStrangerInfo 获取陌生人信息
//// noCatch指定是否使用缓存
//func GetStrangerInfo(qq int64, noCatch bool) string {
//	callAPI(apiPayload{
//		Action: "get_stranger_info",
//		Params: apiParams{},
//	},)
//}

// SendPrivateMsg 发送私聊消息
func SendPrivateMsg(qq int64, msg string) int32 {
	return callAPI(apiPayload{
		Action: "send_private_msg",
		Params: apiParams{"user_id": qq, "message": msg},
	}, nil)
}

// SendGroupMsg 发送群聊消息
func SendGroupMsg(group int64, msg string) int32 {
	return callAPI(apiPayload{
		Action: "send_group_msg",
		Params: apiParams{"group_id": group, "message": msg},
	}, nil)
}
