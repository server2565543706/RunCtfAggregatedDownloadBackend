package Db

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddContent(name string, download string, notes string) {
	var informationTable InformationTable
	if notes == "" {
		informationTable = InformationTable{Name: name, Connect: download, Notes: "无"}
	} else {
		// 在这里可以使用 main.Engine 来访问数据库连接引擎
		informationTable = InformationTable{Name: name, Connect: download, Notes: notes}
	}

	if Engine == nil {
		log.Println("数据库引擎未初始化")
		return
	}

	n, err := Engine.Insert(&informationTable)
	if err != nil {
		log.Println("插入数据时发生错误:", err)
		return
	}
	if n == 1 {
		log.Println("插入数据成功")
	}
}
func AddAdministratorInformation(name string, username string, password string, email string, token string, access string) {
	administratorinformation := AdministratorInformation{Name: name, Username: username, Password: password, Email: email, Token: token, Access: access}
	if Engine == nil {
		log.Println("数据库引擎未初始化")
		return
	}
	n, err := Engine.Insert(&administratorinformation)
	if err != nil {
		log.Println("插入数据时发生错误:", err)
		return
	}
	if n == 1 {
		log.Println("插入数据成功")
	}

}
