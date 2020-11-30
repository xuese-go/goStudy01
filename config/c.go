package config

import (
	"log"
	"os"
)

var C Cfg

func init() {
	if e := C.getCfg(); e != nil {
		log.Panicln("配置文件加载错误")
		os.Exit(0)
	}
}
