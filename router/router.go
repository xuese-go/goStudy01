/**
整体路由
*/
package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	alcoholRouter "github.com/xuese-go/goStudy01/api/alcohol/router"
	brandRouter "github.com/xuese-go/goStudy01/api/brand/router"
	"github.com/xuese-go/goStudy01/api/file/controller"
	JurUserController "github.com/xuese-go/goStudy01/api/jurUser/controller"
	JurController "github.com/xuese-go/goStudy01/api/jurisdiction/controller"
	loginController "github.com/xuese-go/goStudy01/api/login/controller"
	noticeController "github.com/xuese-go/goStudy01/api/notice/controller"
	"github.com/xuese-go/goStudy01/api/respone/structs"
	seriesRouter "github.com/xuese-go/goStudy01/api/series/router"
	userController "github.com/xuese-go/goStudy01/api/user/controller"
	"github.com/xuese-go/goStudy01/api/user/service"
	"github.com/xuese-go/goStudy01/cache"
	"github.com/xuese-go/goStudy01/config"
	"github.com/xuese-go/goStudy01/log"
	"github.com/xuese-go/goStudy01/util/ip"
	"github.com/xuese-go/goStudy01/util/jwt"
	"github.com/xuese-go/goStudy01/util/md5"
	"net/http"
	"strings"
	"time"
)

/*
初始化
*/
func init() {
	// 加载默认配置
	r := gin.Default()

	//相关配置
	options(r)

	// 路由
	routers(r)

	// 启动并监听默认8080端口
	_ = r.Run(":8080")
}

//路由相关配置 Engine
func options(r *gin.Engine) {

	//生产中使用释放模式,该模式会忽略一些输出,例如加载的静态资源和注册的路由地址等
	if config.C.Dev > 1 {
		gin.SetMode(gin.ReleaseMode)
	}

	/**
	以下为 引擎 engine 默认值
	// New returns a new blank Engine instance without any middleware attached.
	//New返回一个新的空白引擎实例，没有附加任何中间件。
	// By default the configuration is:
	//默认情况下，配置为：
	//如果当前路由无法匹配但
	//存在带有（不带）尾部斜杠的路径的处理程序。
	//例如，如果请求了/foo/但路由只存在于/foo，则
	//对于GET请求，客户端重定向到/foo，http状态代码为301
	//307用于所有其他请求方法。
	// - RedirectTrailingSlash:  true
	//如果启用，路由器将尝试修复当前请求路径（如果没有）
	//已为其注册句柄。
	//第一个多余的路径元素，如../或//被删除。
	//之后，路由器对清理后的路径进行不区分大小写的查找。
	//如果可以找到此路由的句柄，路由器将进行重定向
	//对于GET请求，状态代码为301，对于GET请求，状态代码为307
	//所有其他请求方法。
	//例如/FOO和/。//FOO可以重定向到/FOO。
	//RedirectTrailingSlash独立于此选项。
	// - RedirectFixedPath:      false
	//如果启用，路由器将检查是否允许另一种方法用于
	//当前路由，如果当前请求无法路由。
	//如果是这种情况，则用“方法不允许”来回答请求
	//和HTTP状态码405。
	//如果不允许使用其他方法，则将请求委托给NotFound
	//处理程序。
	//猜测这样可能就不会出现404
	// - HandleMethodNotAllowed: false
	//由客户端IP转发
	//猜测跟重定向方式相关
	// - ForwardedByClientIP:    true
	//如果true， url.RawPath 会被用来查找参数
	// - UseRawPath:             false
	//如果为true，则将取消对路径值的转义。
	//如果UseRawPath为false（默认情况下），则UnescapePathValues为true，
	//作为url.Path路径将被使用，这已经是不可替代的。
	// - UnescapePathValues:     true
	//请求的最大内存 默认32MB
	MaxMultipartMemory int64
	*/

	//logger, _ := zap.NewProduction()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(log.SugarLogger.Desugar(), time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(log.SugarLogger.Desugar(), true))
}

// 路由绑定路径集合
func routers(r *gin.Engine) {

	//模板路径-html文件地址
	r.LoadHTMLGlob("views/**/*")
	//静态文件路径
	r.Static("/static", "static")
	//上传图片使用
	r.Static("/file", config.C.File.Path)

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
		//	聊天
		ind2.GET("/chat", func(context *gin.Context) {
			context.HTML(http.StatusOK, "chat/chat.html", nil)
		})
		//	权限管理
		ind2.GET("/jurisdiction", func(context *gin.Context) {
			context.HTML(http.StatusOK, "jurisdiction/jurisdiction.html", nil)
		})
		//	用户权限管理
		ind2.GET("/jurUser", func(context *gin.Context) {
			context.HTML(http.StatusOK, "jurUser/jurUser.html", nil)
		})
		//	brand
		brandRouter.BrandRouterHtml(ind2)
		//	series
		seriesRouter.SeriesRouterHtml(ind2)
		//	alcohol
		alcoholRouter.AlcoholRouterHtml(ind2)
	}

	//api路由
	apis := r.Group("/api")
	{

		//跨域
		apis.Use(Cors())
		//接口次数限制
		apis.Use(restrictions())
		//令牌缓存是否存在判断
		apis.Use(cache.CreateContainersFactory().IsCacheToken())
		//令牌合法性判断
		apis.Use(interceptToken())

		//file
		file := apis.Group("/file")
		file.POST("/up", controller.Up)
		file.GET("/dow", controller.Dow)

		//login
		login := apis.Group("/login")
		login.POST("/login", loginController.Login)

		//	notice
		notice := apis.Group("/notice")
		notice.GET("/notice", noticeController.GetNotices)

		//user
		user := apis.Group("/user")
		user.POST("/user", isAdmin(), userController.Save)
		user.DELETE("/user/:id", isAdmin(), userController.Delete)
		user.PUT("/user/:id", isAdmin(), userController.Update)
		user.PUT("/file", userController.UpdateImg)
		user.GET("/user", isAdmin(), userController.Page)
		user.GET("/user/:id", userController.One)
		user.GET("/userInfo", userController.Info)
		user.GET("/rest/pwd/:id", isAdmin(), userController.RestPwd)
		//jurisdiction
		jurisdiction := apis.Group("/jurisdiction")
		jurisdiction.POST("/jurisdiction", isAdmin(), JurController.Save)
		jurisdiction.DELETE("/jurisdiction/:id", isAdmin(), JurController.Delete)
		jurisdiction.PUT("/jurisdiction/:id", isAdmin(), JurController.Update)
		jurisdiction.GET("/jurisdiction/:id", JurController.One)
		jurisdiction.GET("/jurisdiction", isAdmin(), JurController.Page)
		///jurUser
		jurUser := apis.Group("/jurUser")
		jurUser.PUT("/jurUser/:id", isAdmin(), JurUserController.Update)
		jurUser.GET("/jurUser/:id", JurUserController.FindByUserId)
		//	brand
		brandRouter.BrandRouter(apis)
		//	series
		seriesRouter.SeriesRouter(apis)
		//	alcohol
		alcoholRouter.AlcoholRouter(apis)
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
	// 	log.SugarLogger.Errorf(file.Filename)

	// 	// Upload the file to specific dst.
	// 	// c.SaveUploadedFile(file, dst)

	// 	c.String(http.StatusOK, log.Sprintf("'%s' uploaded!", file.Filename))
	// })

}

//token 中间件
func interceptToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//是否是登录
		p := c.Request.URL.Path
		if strings.Contains(p, "api/login/") {
			c.Next()
		} else {
			//	判断token
			token := c.Request.Header.Get("xueSeToken")
			if token != "" {
				//令牌合法性验证
				if b, t := jwt.ParseToken(token); b {
					//获取用户ip
					ip2 := ip.GetIp(c)
					//是否与签发ip一致
					if t.Subject == md5.Enc(ip2, "逗你玩!!!") {
						t2, _ := jwt.GenerateToken(t.Issuer, ip2)
						cache.CreateContainersFactory().AddToken(t2, ip2)
						c.Header("token", t2)
						c.Next()
					} else {
						c.JSON(http.StatusInternalServerError, structs.ResponeStruct{Success: false, Msg: "请从新登录R1", Data: "logout"})
						c.Abort()
					}
				} else {
					c.JSON(http.StatusInternalServerError, structs.ResponeStruct{Success: false, Msg: "请从新登录R2", Data: "logout"})
					c.Abort()
				}
			} else {
				c.JSON(http.StatusInternalServerError, structs.ResponeStruct{Success: false, Msg: "请从新登录R3", Data: "logout"})
				c.Abort()
			}
		}
	}
}

//是否是管理员
func isAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("xueSeToken")
		if b, t := jwt.ParseToken(token); !b {
			context.Abort()
			context.JSON(http.StatusInternalServerError, structs.ResponeStruct{Success: false, Msg: "请从新登录", Data: "logout"})
		} else {
			r := service.IsRole(t.Issuer)
			if !r.Success {
				context.Abort()
				context.JSON(http.StatusInternalServerError, structs.ResponeStruct{Success: false, Msg: "该账号不是管理员", Data: "!admin"})
			}
		}
	}
}
