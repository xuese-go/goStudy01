package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xuese-go/goStudy01/api/brand/service"
	"github.com/xuese-go/goStudy01/api/brand/structs"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"log"
	"strconv"
)

/*
新增
*/
func Save(ctx *gin.Context) {
	var mod structs.BrandStructs
	if err := ctx.ShouldBind(&mod); err != nil {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.Panic(err.Error())
		return
	}

	resp.Respone(ctx, service.Save(mod))
}

/*
根据主键删除
*/
func Delete(ctx *gin.Context) {
	uuid := ctx.Param("deleteId")
	respond := service.DeleteById(uuid)
	resp.Respone(ctx, respond)
}

/**
根据id修改
*/
func Update(ctx *gin.Context) {
	uuid := ctx.Param("putId")
	var mod structs.BrandStructs
	if err := ctx.ShouldBind(&mod); err != nil {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.Panic(err.Error())
		return
	}

	mod.Uuid = uuid
	respond := service.Update(mod)
	resp.Respone(ctx, respond)
}

/**
根据主键查询
*/
func One(ctx *gin.Context) {
	uuid := ctx.Param("getId")
	respond := service.One(uuid)
	resp.Respone(ctx, respond)
}

/**
分页
*/
func Page(ctx *gin.Context) {
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	name := ctx.Query("name")

	var mod structs.BrandStructs
	mod.Name = name
	n, _ := strconv.Atoi(pageNum)
	s, _ := strconv.Atoi(pageSize)
	res := service.Page(n, s, mod)
	resp.Respone(ctx, res)
}
