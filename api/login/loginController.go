package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*登录*/
func Login(context *gin.Context) {
	acc := context.PostForm("account")
	pwd := context.PostForm("password")
	ap := "111@qq.com11111"
	if acc+pwd == ap {
		context.JSON(http.StatusOK, gin.H{
			"msg": "成功",
		})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "失败",
		})
	}
}
