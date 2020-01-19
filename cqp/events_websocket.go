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
	log.WithField("url", urlStr).Info("成功连接/event")
	go func() { // 接收数据
		for {
			var e event
			if err := conn.ReadJSON(&e); err != nil {
				log.WithError(err).Fatal("接收Event出错")
			}
			log.WithField("e", e).Debug("事件触发")

			// 根据post_type对事件分类处理
			switch e.PostType {
			case "message":
				switch e.MsgType {
				case "private":
					if PrivateMsg != nil {
						go PrivateMsg(subTypes[e.SubType], e.MsgID, e.UserID, e.RawMsg, e.Font)
					}
				case "group":
					if GroupMsg != nil {
						go GroupMsg(subTypes[e.SubType], e.MsgID, e.GroupID, e.UserID, e.Anonymous.Flag, e.RawMsg, e.Font)
					}
				case "discuss":
					if DiscussMsg != nil {
						// TODO: 不知道这里的subtype应是什么，暂时放个1在这
						go DiscussMsg(1, e.MsgID, e.DiscussID, e.UserID, e.RawMsg, e.Font)
					}
				}

			case "notice":
			case "request":
			default:
				log.WithField("e", e).
					Error("未知的post_type")
			}
		}
	}()
}

var subTypes = map[string]int32{
	"friend":  11,
	"other":   1,
	"group":   2,
	"discuss": 3,

	"normal":    1,
	"anonymous": 1,
	"notice":    1,
}

type event struct {
	MsgID     int32  `json:"message_id"`
	Font      int32  `json:"font"`
	UserID    int64  `json:"user_id"`
	GroupID   int64  `json:"group_id"`
	DiscussID int64  `json:"discuss_id"`
	PostType  string `json:"post_type"`
	MsgType   string `json:"message_type"`
	SubType   string `json:"sub_type"`
	RawMsg    string `json:"raw_message"`
	Anonymous struct {
		Flag string `json:"flag"`
	} `json:"anonymous"`
}
