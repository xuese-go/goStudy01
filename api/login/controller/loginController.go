package controller

import (
	"github.com/gin-gonic/gin"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/api/user/service"
	"github.com/xuese-go/goStudy01/api/user/structs"
)

/*登录*/
func Login(context *gin.Context) {
	acc := context.PostForm("account")
	pwd := context.PostForm("password")
	r := service.ByAccount(acc)
	if r.Success {
		if r.Data.(structs.UserStruct).Password == pwd {
			resp.Respone(context, resp.ResponeStruct{Success: true, Data: r.Data.(structs.UserStruct).Uuid})
		}
	} else {
		resp.Respone(context, r)
	}
}
