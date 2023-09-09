package api

import (
	"GoWeb/src/Db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// AddAdministratorInformation 注册管理员账号
func AddAdministratorInformation(context *gin.Context) {
	var data AdministratorInformation
	if err := context.BindJSON(&data); err != nil {
		// 处理错误，例如返回错误响应
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.Name != "" && data.Username != "" && data.Password != "" && data.Email != "" && data.Access != "" {
		getname, getusername, err := Db.GetRepeatedInformation(data.Name, data.Username)
		if getname != "" && getusername != "" {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": "名称或账号重复",
			})
		} else if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err,
			})
		} else if err == nil {
			Token := AddToken(data.Name)
			Db.AddAdministratorInformation(data.Name, data.Username, data.Password, data.Email, Token, data.Access)
			context.JSON(http.StatusOK, gin.H{
				"data": "注册成功",
			})
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "停止你的渗透测试!",
		})
	}
}

// GetAdministratorInformation 登录管理员后台
func GetAdministratorInformation(context *gin.Context) {
	var data AdministratorInformation
	if err := context.BindJSON(&data); err != nil {
		// 处理错误，例如返回错误响应
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.Username != "" && data.Password != "" {
		Token, Name, Access, err := Db.GetAdministratorInformation(data.Username, data.Password)
		if err != nil {
			log.Println("error", err)
		} else if Token != "" {
			context.JSON(http.StatusOK, gin.H{
				"Token":  Token,
				"Name":   Name,
				"Access": Access,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": "账号密码错误",
			})
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "停止你的渗透测试",
		})
	}
}

func GetTokenInformation(context *gin.Context) {
	//token, err := context.Cookie("token") // 从 Cookie 中获取 token
	//if err != nil {
	//	log.Println(err)
	//}
	//getname, erro := Db.GetLogToken(token)
	//if erro != nil {
	//	log.Println(erro)
	//} else {
	//	context.JSON(http.StatusOK, gin.H{
	//		"name": getname,
	//	})
	//}
}
