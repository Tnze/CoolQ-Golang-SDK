// +build websocket

package cqp

import (
	"encoding/json"
	"net/http"

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

	// 接收数据
	for {
		var e apiRet
		if err := conn.ReadJSON(&e); err != nil {
			log.WithError(err).Fatal("API出错")
		}
		log.Info(e)
	}
}

type apiRet struct {
	Status  string          `json:"status"`
	RetCode int             `json:"retcode"`
	Echo    int             `json:"echo"`
	Data    json.RawMessage `json:"data"`
}

func SendPrivateMsg(qq int64, msg string) int32 {
	log.WithField("qq", qq).
		WithField("msg", msg).
		Debug("SendPrivateMsg")
	return 0
}
