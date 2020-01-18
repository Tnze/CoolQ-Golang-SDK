// +build websocket

package cqp

import (
	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func connEvent(urlStr string, requestHeader http.Header) {
	conn, _, err := ws.DefaultDialer.Dial(urlStr, requestHeader)
	if err != nil {
		log.WithError(err).Fatal("连接/event失败")
	}
	log.WithField("remoteAddr", urlStr).Info("成功连接/event")

	// 接收数据
	for {
		var e interface{}
		if err := conn.ReadJSON(&e); err != nil {
			log.WithError(err).Fatal("Event出错")
		}
		log.Info(e)
	}

}
