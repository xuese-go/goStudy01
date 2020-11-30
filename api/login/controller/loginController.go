package controller

import (
	"github.com/gin-gonic/gin"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/api/user/service"
	"github.com/xuese-go/goStudy01/api/user/structs"
	"github.com/xuese-go/goStudy01/util/jwt"
	"github.com/xuese-go/goStudy01/util/md5"
)

/*登录*/
func Login(context *gin.Context) {
	acc := context.PostForm("account")
	pwd := context.PostForm("password")
	r := service.ByAccount(acc)
	if r.Success {
		if r.Data.(structs.UserStruct).State == 2 {
			resp.Respone(context, resp.ResponeStruct{Success: false, Msg: "该账号已被停止使用"})
			return
		}
		//密码加密
		pwd = md5.Enc(pwd)
		if r.Data.(structs.UserStruct).Password == pwd {
			//生成token
			if token, err := jwt.GenerateToken(r.Data.(structs.UserStruct).Uuid); err != nil {
				resp.Respone(context, resp.ResponeStruct{Success: false, Msg: "令牌生成失败"})
				return
			} else {
				resp.Respone(context, resp.ResponeStruct{Success: true, Data: token})
				return
			}
		} else {
			resp.Respone(context, resp.ResponeStruct{Success: false, Msg: "账号或密码错误"})
			return
		}
	} else {
		resp.Respone(context, r)
	}
}
