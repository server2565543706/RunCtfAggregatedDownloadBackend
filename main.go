package main

import (
	"GoWeb/src/Db"
	"GoWeb/src/Router"
)

func main() {
	// 数据库初始化
	Db.GetMysql()
	// Web初始化
	WebServer := Router.Router()
	err := WebServer.Run(":8081")
	if err != nil {
		return
	}
}
