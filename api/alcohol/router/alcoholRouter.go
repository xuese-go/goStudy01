/**

 */
package router

import (
	"github.com/gin-gonic/gin"
	alcoholRouter "github.com/xuese-go/goStudy01/api/alcohol/controller"
	"net/http"
)

/**
页面路由
*/
func AlcoholRouterHtml(r *gin.RouterGroup) {
	r.GET("/alcohol", func(context *gin.Context) {
		context.HTML(http.StatusOK, "alcohol/alcohol.html", nil)
	})
}

/**
api路由
*/
func AlcoholRouter(r *gin.RouterGroup) {
	alcohol := r.Group("/alcohol")
	alcohol.POST("/alcohol", alcoholRouter.Save)
	alcohol.DELETE("/alcohol/:id", alcoholRouter.Delete)
	alcohol.PUT("/alcohol/:id", alcoholRouter.Update)
	alcohol.GET("/alcohol/:id", alcoholRouter.One)
	alcohol.GET("/alcohol", alcoholRouter.Page)
}
