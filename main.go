package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载默认配置
	r := gin.Default()

	// // 路由
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

	// 启动并监听默认8080端口
	r.Run()
}
