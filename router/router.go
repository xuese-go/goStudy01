package router

import (
	"github.com/gin-gonic/gin"
	loginController "github.com/xuese-go/goStudy01/api/login/controller"
	noticeController "github.com/xuese-go/goStudy01/api/notice/controller"
	"github.com/xuese-go/goStudy01/api/respone/structs"
	userController "github.com/xuese-go/goStudy01/api/user/controller"
	"github.com/xuese-go/goStudy01/api/user/service"
	"github.com/xuese-go/goStudy01/api/util/jwt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
初始化
*/
func init() {
	// 加载默认配置
	r := gin.Default()

	//日志
	file, _ := os.Create("goStudy01.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

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
			context.HTML(http.StatusOK, "index/index.html", nil)
		})

		ind2 := index.Group("/page")
		//home页面
		ind2.GET("/home", func(context *gin.Context) {
			context.HTML(http.StatusOK, "home/home.html", nil)
		})
		//	用户管理页面
		ind2.GET("/user", func(context *gin.Context) {
			context.HTML(http.StatusOK, "user/user.html", nil)
		})
		//	个人资料页面
		ind2.GET("/user/personal", func(context *gin.Context) {
			context.HTML(http.StatusOK, "user/personal.html", nil)
		})
	}

	//api路由
	apis := r.Group("/api", interceptToken())
	{
		//login
		login := apis.Group("/login")
		login.POST("/login", loginController.Login)

		//	notice
		notice := apis.Group("/notice")
		notice.GET("/notice", noticeController.GetNotices)

		//user
		user := apis.Group("/user")
		user.POST("/user", isAdmin(), userController.Save)
		user.DELETE("/user/:deleteId", isAdmin(), userController.Delete)
		user.PUT("/user/:putId", isAdmin(), userController.Update)
		user.GET("/user/:getId", userController.One)
		user.GET("/users", isAdmin(), userController.Page)
		user.GET("/user", userController.Info)
		user.GET("/rest/pwd/:restId", isAdmin(), userController.RestPwd)
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

	// 	c.String(http.StatusOK, log.Sprintf("'%s' uploaded!", file.Filename))
	// })

}

//token 中间件
func interceptToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//是否是登录
		path := c.Request.URL.Path
		if strings.Contains(path, "api/login/") {
			c.Next()
		} else {
			//	判断token
			token := c.Request.Header.Get("xueSeToken")
			log.Println("token:", token)
			if token != "" {
				//获取token中的信息
				if claims, err := jwt.ParseToken(token); err != nil {
					c.Abort()
					c.JSON(http.StatusUnauthorized, structs.ResponeStruct{Success: false, Msg: "令牌解析失败", Data: "logout"})
					return
				} else {
					//判断令牌合法性
					if jwt.IsToken(token) {
						//	令牌非法或过期
						c.Abort()
						c.JSON(http.StatusUnauthorized, structs.ResponeStruct{Success: false, Msg: "令牌过期，请从新登录", Data: "logout"})
						return
					} else {
						//确认账号状态
						r := service.IsState(claims.Uuid)
						if r.Success {
							c.Abort()
							c.JSON(http.StatusUnauthorized, structs.ResponeStruct{Success: false, Msg: "该账号已被停用或删除", Data: "logout"})
							return
						} else {
							c.Next() //写不写 都会执行
						}
					}
				}
			} else {
				c.Abort()
				c.JSON(http.StatusUnauthorized, structs.ResponeStruct{Success: false, Msg: "请从新登录", Data: "logout"})
				return
			}
		}
	}
}

//是否是管理员
func isAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("xueSeToken")
		if claims, err := jwt.ParseToken(token); err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, structs.ResponeStruct{Success: false, Msg: "令牌解析错误"})
		} else {
			r := service.IsRole(claims.Uuid)
			if !r.Success {
				context.Abort()
				context.JSON(http.StatusUnauthorized, structs.ResponeStruct{Success: false, Msg: "该账号不是管理员", Data: "!admin"})
			}
		}
	}
}
