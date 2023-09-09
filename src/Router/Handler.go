package Router

import (
	"GoWeb/src/api"
	"github.com/gin-gonic/gin"
)

func myHandler() gin.HandlerFunc {
	return api.GetToken
}
func myHandlerGetToken() gin.HandlerFunc {
	return api.MyHandlerGetToken
}
