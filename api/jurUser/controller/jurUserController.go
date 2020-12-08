package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xuese-go/goStudy01/api/jurUser/service"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"strings"
)

/**
修改指定人权限
*/
func Update(ctx *gin.Context) {
	uuid := ctx.Param("userId")
	jurStr := ctx.PostForm("jurStr")
	jurs := make([]string, 0)
	if jurStr != "" {
		jurs = strings.Split(jurStr, ",")
	}
	respond := service.Update(uuid, jurs)
	resp.Respone(ctx, respond)
}

/**
根据用户主键查询
*/
func FindByUserId(ctx *gin.Context) {
	uuid := ctx.Param("userId")
	respond := service.FindByUserId(uuid)
	resp.Respone(ctx, respond)
}
