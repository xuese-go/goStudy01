package controller

import (
	"github.com/gin-gonic/gin"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/api/user/service"
	"github.com/xuese-go/goStudy01/api/user/structs"
	"github.com/xuese-go/goStudy01/log"
	"github.com/xuese-go/goStudy01/util/file"
	"github.com/xuese-go/goStudy01/util/jwt"
	"strconv"
)

/*
新增用户
*/
func Save(ctx *gin.Context) {
	pwd2 := ctx.PostForm("password2")
	var user structs.UserStruct
	if err := ctx.ShouldBind(&user); err != nil {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.SugarLogger.Errorf(err.Error())
		return
	}

	if user.Password == pwd2 {
		respond := service.Save(user)
		resp.Respone(ctx, respond)
	} else {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "两次密码不一致"})
	}
}

/*
根据主键删除
*/
func Delete(ctx *gin.Context) {
	uuid := ctx.Param("id")
	respond := service.DeleteById(uuid)
	resp.Respone(ctx, respond)
}

/**
根据id修改
*/
func Update(ctx *gin.Context) {
	uuid := ctx.Param("id")
	var user structs.UserStruct
	if err := ctx.ShouldBind(&user); err != nil {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.SugarLogger.Errorf(err.Error())
		return
	}

	user.Uuid = uuid
	respond := service.Update(user)
	resp.Respone(ctx, respond)
}

/**
根据id修改
*/
func UpdateImg(context *gin.Context) {
	token := context.Request.Header.Get("xueSeToken")
	if b, t := jwt.ParseToken(token); !b {
		resp.Respone(context, resp.ResponeStruct{Success: false, Msg: "请从新登录", Data: "logout"})
	} else {
		if str := file.Up(context); str != "" {
			respond := service.UpdateImg(str, t.Issuer)
			resp.Respone(context, respond)
		} else {
			resp.Respone(context, resp.ResponeStruct{Success: false, Msg: "上传失败"})
		}
	}
}

/**
根据id修改
*/
func RestPwd(ctx *gin.Context) {
	uuid := ctx.Param("id")
	var user structs.UserStruct
	if err := ctx.ShouldBind(&user); err != nil {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.SugarLogger.Errorf(err.Error())
		return
	}

	user.Uuid = uuid
	user.Password = "rest"
	respond := service.Update(user)
	resp.Respone(ctx, respond)
}

/**
根据主键查询
*/
func One(ctx *gin.Context) {
	uuid := ctx.Param("id")
	respond := service.One(uuid)
	resp.Respone(ctx, respond)
}

/**
根据令牌查询
*/
func Info(ctx *gin.Context) {
	token := ctx.Request.Header.Get("xueSeToken")
	if b, t := jwt.ParseToken(token); !b {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "请从新登录", Data: "logout"})
	} else {
		respond := service.One(t.Issuer)
		resp.Respone(ctx, respond)
	}
}

/**
分页
*/
func Page(ctx *gin.Context) {
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	acc := ctx.Query("account")

	var user structs.UserStruct
	user.Account = acc
	n, _ := strconv.Atoi(pageNum)
	s, _ := strconv.Atoi(pageSize)
	res := service.Page(n, s, user)

	resp.Respone(ctx, res)
}
