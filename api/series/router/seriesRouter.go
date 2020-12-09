/**

 */
package router

import (
	"github.com/gin-gonic/gin"
	seriesRouter "github.com/xuese-go/goStudy01/api/series/controller"
	"net/http"
)

/**
页面路由
*/
func SeriesRouterHtml(r *gin.RouterGroup) {
	r.GET("/series", func(context *gin.Context) {
		context.HTML(http.StatusOK, "series/series.html", nil)
	})
}

/**
api路由
*/
func SeriesRouter(r *gin.RouterGroup) {
	series := r.Group("/series")
	series.POST("/series", seriesRouter.Save)
	series.DELETE("/series/:deleteId", seriesRouter.Delete)
	series.PUT("/series/:putId", seriesRouter.Update)
	series.GET("/series/:getId", seriesRouter.One)
	series.GET("/seriess", seriesRouter.Page)
}
