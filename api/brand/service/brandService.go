package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xuese-go/goStudy01/api/brand/structs"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/db"
	"github.com/xuese-go/goStudy01/log"
	util "github.com/xuese-go/goStudy01/util/page"
	"time"
)

/**
新增
*/
func Save(mod structs.BrandStructs) resp.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()

	//查询是否有重复
	var u structs.BrandStructs
	_ = dba.First(&u, "name = ?", mod.Name)
	if u.Uuid != "" {
		return resp.ResponeStruct{Success: false, Msg: "重复"}
	}

	//填充其它数据
	uid := uuid.NewV4().String()
	mod.Uuid = uid
	mod.CreateTime = time.Now()

	//新增数据
	t := tx.Create(mod)
	if t.Error != nil {
		t.Rollback()
		return resp.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	t.Commit()
	return resp.ResponeStruct{Success: true, Msg: "操作成功"}
}

/**
删除
*/
func DeleteById(uuid string) resp.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()
	var u structs.BrandStructs
	if err := tx.First(&u, "uuid = ?", uuid).Delete(&u).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		tx.Rollback()
		return resp.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	tx.Commit()
	return resp.ResponeStruct{Success: true, Msg: "操作成功"}
}

/**
修改
*/
func Update(mod structs.BrandStructs) resp.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()
	var u structs.BrandStructs
	if err := dba.First(&u, "uuid = ?", mod.Uuid).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "查询错误"}
	}

	// 需要修改的字段
	if mod.Name != "" {
		u.Name = mod.Name
	}

	u.LastUpdateTime = time.Now()
	t := tx.Save(&u)
	if t.Error != nil {
		t.Rollback()
		log.SugarLogger.Errorf(t.Error.Error())
		return resp.ResponeStruct{Success: false, Msg: "失败"}
	}
	t.Commit()
	return resp.ResponeStruct{Success: true, Msg: "成功"}
}

/**
根据id查询
*/
func One(uuid string) resp.ResponeStruct {
	dba := db.Db
	var u structs.BrandStructs
	if err := dba.First(&u, "uuid = ?", uuid).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "查询错误"}
	}
	return resp.ResponeStruct{Success: true, Msg: "操作成功", Data: u}
}

/**
分页查询
*/
func Page(pageNum int, pageSize int, mod structs.BrandStructs) resp.ResponeStruct {
	//为了不影响后边的操作  所以需要使用新的变量
	dba := db.Db
	us := make([]structs.BrandStructs, 0)

	//查询条件
	if mod.Name != "" {
		dba = dba.Where("name like ?", "%"+mod.Name+"%")
	}

	//总记录数
	var count int
	if pageNum > 0 && pageSize > 0 {
		dba = dba.Find(&us).Count(&count)
		if dba.Error != nil {
			log.SugarLogger.Errorf(dba.Error.Error())
			return resp.ResponeStruct{Success: false, Msg: "操作失败"}
		}

		//分页信息
		dba = dba.Order("name")
		dba = dba.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	//查询
	if err := dba.Table("brand_table").Select([]string{"uuid", "name"}).Scan(&us).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	if pageNum > 0 && pageSize > 0 {

		return resp.ResponeStruct{Success: true, Data: us, Page: util.PageUtil(count, pageSize, pageNum)}
	} else {
		return resp.ResponeStruct{Success: true, Data: us}
	}
}
