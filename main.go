package main

import (
	"login-test/api"
	"login-test/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
}
