package api

import (
	"GoWeb/src/Db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// AddToken 生成Token
func AddToken(username string) string {
	// 创建一个新的 Token 对象
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置 Token 中的声明信息
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 设置过期时间为 24 小时

	// 使用密钥对 Token 进行签名
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		// 处理生成 Token 失败的情况
		// 例如记录日志、返回错误信息等
	}

	return tokenString
}

// GetToken 令牌验证
func GetToken(context *gin.Context) {
	token, err := context.Cookie("token") // 从 Cookie 中获取 token
	log.Println("111111+++++>>>", token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":  "未提供有效的令牌",
			"data": nil,
		})
		return
	}

	b, _, erro := Db.GetLogToken(token)
	if erro != nil {
		log.Println("error", erro)
	}
	if b == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":  "无效的令牌",
			"data": nil,
		})
		return
	}

	context.Next()
}

func MyHandlerGetToken(context *gin.Context) {
	var data AdministratorInformation
	if err := context.BindJSON(&data); err != nil {
		// 处理错误，例如返回错误响应
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("111111+++++>>>", data.Token)

	b, access, erro := Db.GetLogToken(data.Token)
	if erro != nil {
		log.Println("error", erro)
	}
	log.Println(time.Time{}, "欢迎用户", b, "自动登录成功")
	if b == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":  "无效的令牌",
			"data": nil,
		})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{
			"name":   b,
			"access": access,
		})
	}

	context.Next()
}
