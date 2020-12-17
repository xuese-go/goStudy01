package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xuese-go/goStudy01/api/jurisdiction/service"
	"github.com/xuese-go/goStudy01/api/jurisdiction/structs"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/log"
	"strconv"
)

/*
新增
*/
func Save(ctx *gin.Context) {
	var jur structs.JurisdictionStruct
	if err := ctx.ShouldBind(&jur); err != nil {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.SugarLogger.Errorf(err.Error())
		return
	}

	resp.Respone(ctx, service.Save(jur))
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
	var jur structs.JurisdictionStruct
	if err := ctx.ShouldBind(&jur); err != nil {
		resp.Respone(ctx, resp.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.SugarLogger.Errorf(err.Error())
		return
	}

	jur.Uuid = uuid
	respond := service.Update(jur)
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
	acc := ctx.Query("jurName")

	if pageNum != "" && pageSize != "" && acc != "" {
		var jur structs.JurisdictionStruct
		jur.JurName = acc
		n, _ := strconv.Atoi(pageNum)
		s, _ := strconv.Atoi(pageSize)
		res := service.Page(n, s, jur)
		resp.Respone(ctx, res)
	} else {
		var jur structs.JurisdictionStruct
		res := service.Page(-1, -1, jur)
		resp.Respone(ctx, res)
	}
}
