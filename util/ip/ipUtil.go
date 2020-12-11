package ip

import (
	"github.com/gin-gonic/gin"
	"log"
)

//获取用户真实ip
func GetIp(c *gin.Context) string {
	ip := c.ClientIP()
	//req := c.Request
	//remoteAddr := req.RemoteAddr
	log.Println(ip)
	//if ip := req.Header.Get(XRealIP); ip != "" {
	//	remoteAddr = ip
	//} else if ip = req.Header.Get(XForwardedFor); ip != "" {
	//	remoteAddr = ip
	//} else {
	//	remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	//}
	//
	//if remoteAddr == "::1" {
	//	remoteAddr = "127.0.0.1"
	//}
	return ip
}
