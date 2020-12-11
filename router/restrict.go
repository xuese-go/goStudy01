/**
限制接口频率
限制频率为同一ip在1s内同一接口只允许调用1次
*/
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/util/ip"
	"net/http"
	"time"
)

var restricts = make(map[string]restriction, 0)

type restriction struct {
	//	请求地址
	p string
	//	请求ip
	i string
	//	请求时间
	t int64
}

func restrictions() gin.HandlerFunc {
	return func(context *gin.Context) {
		//请求类型
		m := context.Request.Method
		//只有 post  put  delete 才会限制
		if m == http.MethodPost || m == http.MethodPut || m == http.MethodDelete {
			//	获取请求地址
			p := context.Request.URL.Path
			//	获取请求ip
			i := ip.GetIp(context)
			//	当前时间(s)
			t := time.Now().Unix()
			if _, ok := restricts[i]; ok {
				//	如果存在
				if p == restricts[i].p {
					//	如果本次请求和记录的上次请求一致
					if (restricts[i].t + 1) < t {
						rs := restriction{
							p: p,
							i: i,
							t: t,
						}
						restricts[i] = rs
						context.Next()
					} else {
						context.Abort()
						context.JSON(http.StatusInternalServerError, structs.ResponeStruct{Success: false, Msg: "请求太频繁"})
					}
				} else {
					rs := restriction{
						p: p,
						i: i,
						t: t,
					}
					restricts[i] = rs
					context.Next()
				}
			} else {
				rs := restriction{
					p: p,
					i: i,
					t: t,
				}
				restricts[i] = rs
				context.Next()
			}
		}
	}
}
