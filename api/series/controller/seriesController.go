/**

 */
package controller

import (
	"github.com/gin-gonic/gin"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/api/series/service"
	"github.com/xuese-go/goStudy01/api/series/structs"
	"github.com/xuese-go/goStudy01/log"
	"strconv"
)

/*
新增
*/
func Save(ctx *gin.Context) {
	var mod structs.SeriesStructs
	if err := ctx.ShouldBind(&mod); err != nil {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.SugarLogger.Errorf(err.Error())
		return
	}

	resp.Respone(ctx, service.Save(mod))
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
	var mod structs.SeriesStructs
	if err := ctx.ShouldBind(&mod); err != nil {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.SugarLogger.Errorf(err.Error())
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
	uuid := ctx.Param("id")
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

	var mod structs.SeriesStructs
	mod.Name = name
	n, _ := strconv.Atoi(pageNum)
	s, _ := strconv.Atoi(pageSize)
	res := service.Page(n, s, mod)
	resp.Respone(ctx, res)
}
