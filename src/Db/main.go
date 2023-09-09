package Db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	Engine *xorm.Engine
)

func GetMysql() {
	var (
		// 数据库连接配置
		userName       string = "root"
		password       string = "123456"
		ipAddress      string = "localhost"
		port           int    = 3306
		dbName         string = "RunCtf"
		charset        string = "utf8mb4"
		dataSourceName        = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	)
	var err error
	// 创建数据库连接引擎
	Engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	// 同步 InformationTable 表结构
	err = Engine.Sync(new(InformationTable))
	// 同步 AdministratorInformation 表结构
	err = Engine.Sync(new(AdministratorInformation))
	if err != nil {
		fmt.Println("表结构同步失败", err)
		fmt.Println("dataSourceName", dataSourceName)
	}
}
