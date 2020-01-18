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

// SendPrivateMsg 发送私聊消息
func SendPrivateMsg(qq int64, msg string) int32 {
	conn := getAPI()
	defer retAPI()

	echo := rand.Int()

	err := conn.WriteJSON(apiPayload{
		Action: "send_private_msg",
		Params: struct {
			UserID  int64  `json:"user_id"`
			Message string `json:"message"`
		}{UserID: qq, Message: msg},
		Echo: echo,
	})
	if err != nil {
		log.WithError(err).
			Error("发送私聊消息失败")
	}

	log.WithField("qq", qq).
		WithField("msg", msg).
		WithField("echo", echo).
		Debug("SendPrivateMsg")

	var resp apiRet
	if err := conn.ReadJSON(&resp); err != nil {
		log.WithError(err).Error("接收私聊消息响应失败")
	}
	log.WithField("echo", resp.Echo).Debug("响应")
	switch resp.Status {
	case "ok":
	case "async":
		log.Warning("请求异步处理，具体成功或失败将无法获知。")
	case "failed":
		log.WithField("retcode", resp.RetCode).
			Error("发送私聊消息失败")
	}
	return 0
}
