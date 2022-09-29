package main

import (
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",  // 输出自定义错误提示
			zap.String("url", url),  // 有关该错误的关键信息
			zap.Error(err))			// 有关该错误的关键信息
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	logger, _ = zap.NewProduction()
    defer logger.Sync()
    simpleHttpGet("www.baidu.com")
}
