package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/api/user/service"
	"github.com/xuese-go/goStudy01/api/user/structs"
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
		return
	}

	if user.Password == pwd2 {
		respond := service.Save(user)
		fmt.Println(respond)
		resp.Respone(ctx, respond)
	} else {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "两次密码不一致"})
	}
}

/*
根据主键删除
*/
func Delete(ctx *gin.Context) {
	uuid := ctx.Param("deleteId")

	resp.Respone(ctx, resp.ResponeStruct{Success: true, Data: uuid})
}

/**
根据id修改
*/
func Update(ctx *gin.Context) {
	uuid := ctx.Param("putId")
	var user structs.UserStruct
	err := ctx.Bind(&user)
	if err != nil {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		return
	}

	user.Uuid = uuid
	resp.Respone(ctx, resp.ResponeStruct{Success: true, Data: user})
}

/**
根据主键查询
*/
func One(ctx *gin.Context) {
	uuid := ctx.Param("getId")

	resp.Respone(ctx, resp.ResponeStruct{Success: true, Data: uuid})
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