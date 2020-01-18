// +build websocket

package cqp

import (
	"flag"
	"net/http"
	"net/url"
	"path"

	"github.com/mattn/go-colorable"

	log "github.com/sirupsen/logrus"
)

var accessToken = flag.String("access_token", "", "API 访问 token")
var serverURL = flag.String("url", "ws://[::]:6700", "Websocket服务器URL")

func init() {
	// 在win环境下也输出好看的日志
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
}

func Main() {
	flag.Parse()
	log.SetLevel(log.DebugLevel)

	// 构造URL
	baseURL, err := url.Parse(*serverURL)
	if err != nil {
		log.WithField("arg_url", *serverURL).
			WithError(err).
			Fatal("解析URL错误")
	}
	var apiURL, eventURL = *baseURL, *baseURL
	apiURL.Path = path.Join(apiURL.Path, "api")
	eventURL.Path = path.Join(eventURL.Path, "event")

	// 连接酷Q
	requestHeader := make(http.Header)
	if *accessToken != "" { // 启用 access_token
		requestHeader.Add("Authorization", "Token "+*accessToken)
	}

	go connAPIs(apiURL.String(), requestHeader)
	connEvent(eventURL.String(), requestHeader)
}
