package controller

import (
	"github.com/gin-gonic/gin"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
)

/*
获取消息
*/
func GetNotices(ctx *gin.Context) {
	notices := make([]string, 0)
	notices = append(notices, "消息1", "消息2", "消息3", "消息4", "消息5")
	resp.Respone(ctx, resp.ResponeStruct{Success: true, Data: notices})
}
