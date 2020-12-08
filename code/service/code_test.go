package service

import (
	"fmt"
	"github.com/xuese-go/goStudy01/code/structs"
	"log"
	"os"
	"testing"
	"text/template"
)

func TestCode(t *testing.T) {
	if t, err := template.ParseFiles("./controller.tmpl"); err != nil {
		log.Panicln("读取模板文件错误")
	} else {
		f := createFile("test.go")
		var c = new(structs.CodeStructs)
		c.Path = "path"
		c.Imps = "imps"
		c.Pck = "pck"
		if err = t.Execute(f, c); err != nil {
			log.Println(err)
		} else {
			if f != nil {
				f.Close()
			}
		}

	}
}

func createFile(fileName string) *os.File {
	//创建文件
	f, err := os.Create(fileName)
	//判断是否出错
	if err != nil {
		fmt.Println(err)
	}
	return f
}
