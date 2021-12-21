package main

import (
	"code.meikeland.com/me/forwarding/api/forwarding"
	_ "code.meikeland.com/me/forwarding/docs"
)

// @title API
// @version 1.0
// @BasePath /api
func main() {
	forwarding.Init()
}
