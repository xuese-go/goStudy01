package config

import (
	"github.com/xuese-go/goStudy01/log"
	"os"
)

var C Cfg

func init() {
	if e := C.getCfg(); e != nil {
		log.SugarLogger.Errorf("配置文件加载错误")
		os.Exit(0)
	}
}
