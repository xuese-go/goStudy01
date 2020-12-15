package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	alcoholStructs "github.com/xuese-go/goStudy01/api/alcohol/structs"
	brandStructs "github.com/xuese-go/goStudy01/api/brand/structs"
	jurisdictionStructs "github.com/xuese-go/goStudy01/api/jurisdiction/structs"
	seriesStructs "github.com/xuese-go/goStudy01/api/series/structs"
	userStructs "github.com/xuese-go/goStudy01/api/user/structs"
	"github.com/xuese-go/goStudy01/config"
	"github.com/xuese-go/goStudy01/log"
)

var Db *gorm.DB

func init() {
	var err error
	dsn := config.C.Db.Dsn //"root:root@tcp(127.0.0.1:3306)/goStudy01?charset=utf8&parseTime=true&loc=Local"
	if Db, err = gorm.Open("mysql", dsn); err != nil {
		log.SugarLogger.Errorf("数据库连接失败")
		log.SugarLogger.Errorf(err.Error())
	} else {
		log.SugarLogger.Infof("数据库连接成功")
		//打印sql
		Db.LogMode(true)
		//创建表
		tables := make([]interface{}, 0)
		tables = append(tables, &userStructs.UserStruct{})
		tables = append(tables, &jurisdictionStructs.JurisdictionStruct{})
		tables = append(tables, &brandStructs.BrandStructs{})
		tables = append(tables, &seriesStructs.SeriesStructs{})
		tables = append(tables, &alcoholStructs.AlcoholStructs{})

		for k := range tables {
			Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(tables[k])
		}
	}

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
