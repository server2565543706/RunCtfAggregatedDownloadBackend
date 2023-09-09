package Router

import (
	"GoWeb/src/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CorsMiddleware 解决跨域问题
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有来源访问，也可以设置为特定的域
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 可选，设置预检请求的有效期，单位为秒

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func Router() *gin.Engine {
	WebServer := gin.Default()

	// 应用跨域
	WebServer.Use(CorsMiddleware())
	GetGroup := WebServer.Group("/api")
	{
		GetGroup.GET("/information", api.GetInformation)

	}
	PostGroup := WebServer.Group("/api")
	{
		PostGroup.POST("/addinformation", myHandler(), api.Addinformation)
		PostGroup.POST("/addadministratorinformation", myHandler(), api.AddAdministratorInformation)
		PostGroup.POST("/getadministratorinformation", api.GetAdministratorInformation)
		PostGroup.POST("/gettokeninformation", myHandlerGetToken(), api.GetTokenInformation)
	}

	return WebServer
}
