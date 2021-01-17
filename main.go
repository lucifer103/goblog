package main

import (
	"net/http"

	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	c "goblog/pkg/config"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
