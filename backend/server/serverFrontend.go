package server

import "net/http"

// 托管 React 前端
func ServerFrontend() {
	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs)
}
