// +build websocket

package cqp

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"

	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

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
	RetCode int             `json:"retcode"`
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

// 接收并检查一般的返回值
func checkRet(conn *ws.Conn, echo int) int32 {
	var resp apiRet
	if err := conn.ReadJSON(&resp); err != nil {
		log.WithError(err).Error("接收私聊消息响应失败")
	}
	// TODO: 检查echo值一致性
	switch resp.Status {
	case "ok":
		log.WithField("echo", resp.Echo).Debug("调用成功")
	case "async":
		log.Warning("请求异步处理，具体成功或失败将无法获知。")
	case "failed":
		log.WithField("retcode", resp.RetCode).
			Error("发送私聊消息失败")
	}
	return 0
}

func GetAppDir() string {
	return *appDir
}

// SendPrivateMsg 发送私聊消息
func SendPrivateMsg(qq int64, msg string) int32 {
	conn := getAPI()
	defer retAPI()

	echo := rand.Int()
	log.WithFields(log.Fields{"msg": msg, "qq": qq, "echo": echo}).Debug("发送私聊消息")

	err := conn.WriteJSON(apiPayload{
		Action: "send_private_msg",
		Params: apiParams{"user_id": qq, "message": msg},
		Echo:   echo,
	})
	if err != nil {
		log.WithError(err).Error("发送私聊消息失败")
	}

	// 检查调用是否成功
	checkRet(conn, echo)
	return 0
}

// SendGroupMsg 发送群聊消息
func SendGroupMsg(qq int64, msg string) int32 {
	conn := getAPI()
	defer retAPI()

	echo := rand.Int()
	log.WithFields(log.Fields{"msg": msg, "qq": qq, "echo": echo}).Debug("发送私聊消息")

	err := conn.WriteJSON(apiPayload{
		Action: "send_private_msg",
		Params: apiParams{"user_id": qq, "message": msg},
		Echo:   echo,
	})
	if err != nil {
		log.WithError(err).Error("发送私聊消息失败")
	}

	// 检查调用是否成功
	checkRet(conn, echo)
	return 0
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
