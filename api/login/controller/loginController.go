package controller

import (
	"github.com/gin-gonic/gin"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
)

/*登录*/
func Login(context *gin.Context) {
	acc := context.PostForm("account")
	pwd := context.PostForm("password")
	ap := "111@qq.com11111"
	if acc+pwd == ap {
		resp.Respone(context, resp.ResponeStruct{Success: true, Data: "abcdddsefssfews"})
	} else {
		resp.Respone(context, resp.ResponeStruct{Success: false, Msg: "账号或密码错误"})
	}
}
