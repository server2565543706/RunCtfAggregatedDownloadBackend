package Db

import (
	"log"
	"time"
)

// FindAllContent 查询所有内容
func FindAllContent() ([]InformationTable, error) {
	var information []InformationTable

	if Engine == nil {
		log.Println("数据库引擎未初始化")
		return nil, nil
	}

	err := Engine.Find(&information)
	if err != nil {
		log.Println("查询失败", err)
		return nil, err
	}

	return information, nil
}

// GetAdministratorInformation 验证管理员账户
func GetAdministratorInformation(username string, password string) (string, string, string, error) {
	var administratorinformation AdministratorInformation
	if Engine == nil {
		log.Println("数据库引擎未初始化")
		return "", "", "", nil
	}
	// 查询数据库中是否存在匹配的用户名和密码
	has, err := Engine.Where("username = ? AND password = ?", username, password).Get(&administratorinformation)
	if err != nil {
		log.Println("查询失败", err)
		return "", "", "", err
	}

	if !has {
		log.Println(time.Time{}, "----->账号为", username, "登录时账号密码错误！")
		return "", "", "", nil
	}
	return administratorinformation.Token, administratorinformation.Name, administratorinformation.Access, nil
}

// GetRepeatedInformation 验证注册时是否有相同用户
func GetRepeatedInformation(name string, username string) (string, string, error) {
	var administratorinformation AdministratorInformation
	if Engine == nil {
		log.Println("数据库引擎未初始化")
		return "", "", nil
	}
	// 查询数据库中是否存在匹配的用户名和密码
	has, err := Engine.Where("name = ? AND username = ?", name, username).Get(&administratorinformation)
	if err != nil {
		log.Println("查询失败", err)
		return "", "", err
	}

	if !has {
		log.Println(time.Time{}, "----->账号为", username, "登录时账号密码错误！")
		return "", "", nil
	}
	return administratorinformation.Name, administratorinformation.Username, nil
}

// GetLogToken 校验Token
func GetLogToken(token string) (string, string, error) {
	var administratorinformation AdministratorInformation
	if Engine == nil {
		log.Println("数据库引擎未初始化")
		return "", "", nil
	}
	// 查询数据库中是否存在匹配的用户名和密码
	has, err := Engine.Where("token = ?", token).Get(&administratorinformation)
	if err != nil {
		log.Println("查询失败", err)
		return "", "", err
	}

	if !has {
		log.Println(time.Time{}, "----->Token为", token, "用户调用了token验证接口")
		return "", "", nil
	}
	return administratorinformation.Name, administratorinformation.Access, nil
}
