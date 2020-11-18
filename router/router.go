package router

import (
	"github.com/gin-gonic/gin"
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

	//主入口
	index := r.Group("/")
	{
		//主模块
		ind := index.Group("/")
		ind.GET("/", func(c *gin.Context) {
			c.HTML(200, "index/index.html", nil)
		})
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
