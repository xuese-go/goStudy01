package service

import (
	"fmt"
	"github.com/xuese-go/goStudy01/code/structs"
	"os"
	"strings"
	"text/template"
)

/**
生成模块代码
*/
func AutoGenerationMod(mod string, modCN string) {
	//	go
	autoGenerationGo(mod, modCN)
	//	html
	autoGenerationHtml(mod, modCN)
	//	js
	autoGenerationJs(mod)
}

/**
自动生成js基本文件和代码
*/
func autoGenerationJs(mod string) {
	//html
	codeUtilJs("js", mod)
}

/**
工具类js
*/
func codeUtilJs(category string, mod string) {
	if tem, err := template.ParseFiles("./../tmpl/" + category + ".tmpl"); err != nil {
		fmt.Println("读取模板文件错误")
	} else {
		f := createFile("./../../static/js/"+mod, mod+".js")
		if f == nil {
			fmt.Println("文件创建失败")
			return
		}
		var c = new(structs.CodeStructs)
		c.ModPath = mod
		if err = tem.Execute(f, c); err != nil {
			fmt.Println(err)
		} else {
			defer f.Close()
		}

	}
}

/**
自动生成html基本文件和代码
*/
func autoGenerationHtml(mod string, modCN string) {
	//html
	codeUtilHtml("html", mod, modCN)
}

/**
工具类html
*/
func codeUtilHtml(category string, mod string, modCN string) {
	if tem, err := template.ParseFiles("./../tmpl/" + category + ".tmpl"); err != nil {
		fmt.Println("读取模板文件错误")
	} else {
		f := createFile("./../../views/"+mod, mod+".html")
		if f == nil {
			fmt.Println("文件创建失败")
			return
		}
		var c = new(structs.CodeStructs)
		c.ModPath = mod
		c.ModStruct = strings.Title(mod)
		c.ModChina = modCN
		c.Top = "{{define \"" + mod + "/" + mod + ".html\"}}"
		c.Ends = "{{end}}"
		if err = tem.Execute(f, c); err != nil {
			fmt.Println(err)
		} else {
			defer f.Close()
		}

	}
}

/**
自动生成go基本文件和代码
*/
func autoGenerationGo(mod string, modCN string) {
	//router
	codeUtilGo("router", mod, "")
	//structs
	codeUtilGo("structs", mod, modCN)
	//service
	codeUtilGo("service", mod, "")
	//controller
	codeUtilGo("controller", mod, "")
}

/**
工具类go
*/
func codeUtilGo(category string, mod string, modCN string) {
	if tem, err := template.ParseFiles("./../tmpl/" + category + ".tmpl"); err != nil {
		fmt.Println("读取模板文件错误")
	} else {
		category2 := strings.Title(category)
		f := createFile("./../../api/"+mod+"/"+category, mod+category2+".go")
		if f == nil {
			fmt.Println("文件创建失败")
			return
		}
		var c = new(structs.CodeStructs)
		c.ModPath = mod
		c.ModStruct = strings.Title(mod)
		c.ControllerName = mod + category2
		c.ModChina = modCN
		if err = tem.Execute(f, c); err != nil {
			fmt.Println(err)
		} else {
			defer f.Close()
		}

	}
}

/**
创建文件和目录
*/
func createFile(filePath string, fileName string) *os.File {
	//判断目录是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		//	如果不存在则创建
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	//判断文件是否存在
	if _, err := os.Stat(filePath + "/" + fileName); !os.IsNotExist(err) {
		//	如果存在则删除文件
		_ = os.Remove(filePath + "/" + fileName)
	}
	//创建文件
	f, err := os.Create(filePath + "/" + fileName)
	//判断是否出错
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return f
}
