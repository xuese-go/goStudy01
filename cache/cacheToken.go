/**
token缓存
*/
package cache

import (
	"github.com/gin-gonic/gin"
	"github.com/xuese-go/goStudy01/api/respone/structs"
	"log"
	"net/http"
	"strings"
	"sync"
)

//线程安全
var sMap sync.Map

// 创建一个容器工厂
func CreateContainersFactory() *cacheToken {
	return &cacheToken{}
}

type cacheToken struct {
}

/**
添加token
*/
func (ct *cacheToken) AddToken(token string, ip string) {
	sMap.Store(ip, token)
	log.Println("新增令牌", token)
}

/**
是否存在，如果存在则删除，并返回true，否则false
*/
func (ct *cacheToken) isToken(token string, ip string) bool {
	if v, ok := sMap.Load(ip); ok {
		if v == token {
			sMap.Delete(ip)
			log.Println("销毁令牌", token)
			return true
		}
	}
	return false
}

/**
判断令牌有没有销毁
*/
func (ct *cacheToken) IsCacheToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//是否是登录
		p := c.Request.URL.Path
		if strings.Contains(p, "api/login/") {
			c.Next()
		} else {
			//	判断token
			token := c.Request.Header.Get("xueSeToken")
			if token != "" {
				if ct.isToken(token, c.ClientIP()) {
					c.Next()
				} else {
					c.JSON(http.StatusInternalServerError, structs.ResponeStruct{Success: false, Msg: "请从新登录C1", Data: "logout"})
					c.Abort()
				}
			} else {
				c.JSON(http.StatusInternalServerError, structs.ResponeStruct{Success: false, Msg: "请从新登录C2", Data: "logout"})
				c.Abort()
			}
		}
	}
}
