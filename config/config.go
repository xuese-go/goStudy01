package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Cfg struct {
	Db   db   `yaml:"db"`
	File file `yaml:"file"`
}

type db struct {
	Dsn string `yaml:"dsn"`
}

type file struct {
	Path string `yaml:"path"`
}

/**
读取配置文件，并转成结构体
*/
func (cfg *Cfg) getCfg() *error {
	f, _ := os.Getwd()
	if c, err := ioutil.ReadFile(f + "/conf.yaml"); err != nil {
		log.Panicln(err.Error())
		return &err
	} else if err = yaml.Unmarshal(c, &cfg); err != nil {
		//	二进制转换为字符，如果有错误
		log.Println(err.Error())
		return &err
	} else {
		return nil
	}
}
