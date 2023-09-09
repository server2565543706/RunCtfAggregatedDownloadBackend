package api

import (
	"GoWeb/src/Db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInformation(context *gin.Context) {
	user, err := Db.FindAllContent()
	if err != nil {

	} else {
		context.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}
func Addinformation(context *gin.Context) {
	// 声明一个结构体，用于解析 JSON 数据
	var data YourStruct // 替换 YourStruct 为您期望的数据结构
	// 使用 BindJSON 方法来解析 JSON 数据并将其绑定到 data 结构体
	if err := context.BindJSON(&data); err != nil {
		// 处理错误，例如返回错误响应
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.Name != "" && data.Download != "" {
		Db.AddContent(data.Name, data.Download, data.Notes)
		// 返回成功响应
		context.JSON(http.StatusOK, gin.H{
			"message":  "JSON data received successfully",
			"Name":     data.Name,
			"Download": data.Download,
			"Notes":    data.Notes,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "停止你的渗透测试!",
		})
	}

}
