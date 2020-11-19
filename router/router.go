package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuese-go/goStudy01/api/login/controller"
)

/*
初始化
*/
func Init() {
	// 加载默认配置
	r := gin.Default()

	// 路由
	routers(r)

	// 启动并监听默认8080端口
	_ = r.Run()
}

// 路由绑定路径集合
func routers(r *gin.Engine) {

	//模板路径-html文件地址
	r.LoadHTMLGlob("views/**/*")
	//静态文件路径
	r.Static("/static", "static")

	//页面路由
	index := r.Group("/")
	{
		//页面处理
		ind := index.Group("/")
		//主页面-登录页面
		ind.GET("/", func(context *gin.Context) {
			context.HTML(200, "index/index.html", nil)
		})

		ind2 := index.Group("/page", interceptToken())
		//home页面
		ind2.GET("/home", func(context *gin.Context) {
			context.HTML(200, "home/home.html", nil)
		})
	}

	//api路由
	apis := r.Group("/api")
	{
		login := apis.Group("/login")
		login.POST("/login", controller.Login)
	}

	// r.GET("/ping/:a/:b", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": c.Param("a") + c.Param("b"),
	// 	})
	// })
	// // 上传
	// r.POST("/upload", func(c *gin.Context) {
	// 	// single file
	// 	file, _ := c.FormFile("file")
	// 	log.Println(file.Filename)

	// 	// Upload the file to specific dst.
	// 	// c.SaveUploadedFile(file, dst)

	// 	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// })

}

//token 中间件
func interceptToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		c.Next()
		fmt.Println("2222222222")
	}
}
