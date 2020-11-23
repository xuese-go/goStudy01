package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goStudy01?charset=utf8&parseTime=true&loc=Local"
	var err error
	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("数据库连接成功")
	}

	//打印sql
	Db.LogMode(true)
	//关闭闲置的连接
	//defer Db.Close()

	//tables := make(map[string]interface{}, 0)
	//tables["user_table"] = &structs.UserStruct{}
	//
	//for k, v := range tables {
	//	//检查表是否存在
	//	log.Printf("检查表%v\n", k)
	//	if !Db.HasTable(k) {
	//		log.Printf("表%q不存在\n", k)
	//		//	创建表
	//		log.Printf("创建表%q\n", k)
	//		d := Db.Table(k).CreateTable(v)
	//		if d.HasTable(k) {
	//			log.Printf("表%q创建成功\n", k)
	//		} else {
	//			log.Printf("表%q创建失败\n", k)
	//		}
	//	} else {
	//		log.Printf("表%v存在\n", k)
	//	}
	//	//	自动迁移表(创建表，添加缺少的列和索引,不会改变现有列的类型且不会删除多余的列)
	//	//Db.AutoMigrate(k)
	//}
}
