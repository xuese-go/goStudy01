package config

import (
	"github.com/xuese-go/goStudy01/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type cfg struct {
	Db   db   `yaml:"db"`
	File file `yaml:"file"`
	Dev  int  `yaml:"dev"`
}

type db struct {
	Dsn string `yaml:"dsn"`
}

type file struct {
	Path string `yaml:"path"`
}

var C = cfg{}

func init() {
	f, _ := os.Getwd()
	if c, err := ioutil.ReadFile(f + "/conf.yaml"); err != nil {
		log.SugarLogger.Errorf(err.Error())
	} else if err = yaml.Unmarshal(c, &C); err != nil {
		//	二进制转换为字符，如果有错误
		log.SugarLogger.Errorf(err.Error())
	}
}
