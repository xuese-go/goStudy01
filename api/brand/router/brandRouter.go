package router

import (
	"github.com/gin-gonic/gin"
	brandRouter "github.com/xuese-go/goStudy01/api/brand/controller"
	"net/http"
)

/**
页面路由
*/
func BrandRouterHtml(r *gin.RouterGroup) {
	r.GET("/brand", func(context *gin.Context) {
		context.HTML(http.StatusOK, "brand/brand.html", nil)
	})
}

/**
api路由
*/
func BrandRouter(r *gin.RouterGroup) {
	brand := r.Group("/brand")
	brand.POST("/brand", brandRouter.Save)
	brand.DELETE("/brand/:id", brandRouter.Delete)
	brand.PUT("/brand/:id", brandRouter.Update)
	brand.GET("/brand/:id", brandRouter.One)
	brand.GET("/brand", brandRouter.Page)
}
