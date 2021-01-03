package main

import (
	"net/http"

	"goblog/app/http/middlewares"
	"goblog/bootstrap"
)

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
