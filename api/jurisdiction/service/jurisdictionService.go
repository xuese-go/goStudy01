package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xuese-go/goStudy01/api/jurisdiction/structs"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/db"
	"github.com/xuese-go/goStudy01/log"
	util "github.com/xuese-go/goStudy01/util/page"
	"time"
)

/**
新增
*/
func Save(jur structs.JurisdictionStruct) resp.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()
	//查询是否有重复账号
	var u structs.JurisdictionStruct
	d := dba.First(&u, "jur_name = ?", jur.JurName)
	if u.JurName != "" {
		return resp.ResponeStruct{Success: false, Msg: "权限名重复"}
	}

	//填充其它数据
	uid := uuid.NewV4().String()
	jur.Uuid = uid
	jur.CreateTime = time.Now()

	//新增数据
	t := tx.Create(jur)
	if t.Error != nil {
		t.Rollback()
		log.SugarLogger.Errorf(d.Error.Error())
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
	var u structs.JurisdictionStruct
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
func Update(jur structs.JurisdictionStruct) resp.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()
	var u structs.JurisdictionStruct
	if err := dba.First(&u, "uuid = ?", jur.Uuid).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "查询错误"}
	}
	if jur.JurName != "" {
		u.JurName = jur.JurName
	}
	if jur.JurFlag != "" {
		u.JurFlag = jur.JurFlag
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
	var u structs.JurisdictionStruct
	if err := dba.First(&u, "uuid = ?", uuid).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "查询错误"}
	}
	return resp.ResponeStruct{Success: true, Msg: "操作成功", Data: u}
}

/**
分页查询
*/
func Page(pageNum int, pageSize int, jur structs.JurisdictionStruct) resp.ResponeStruct {
	//为了不影响后边的操作  所以需要使用新的变量
	dba := db.Db
	us := make([]structs.JurisdictionStruct, 0)

	//查询条件
	if jur.JurName != "" {
		dba = dba.Where("jur_name like ?", "%"+jur.JurName+"%")
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
		dba = dba.Order("jur_name")
		dba = dba.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	//查询
	if err := dba.Table("jurisdiction_table").Select([]string{"uuid", "jur_name", "jur_flag"}).Scan(&us).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	if pageNum > 0 && pageSize > 0 {
		return resp.ResponeStruct{Success: true, Data: us, Page: util.PageUtil(count, pageSize, pageNum)}
	} else {
		return resp.ResponeStruct{Success: true, Data: us}
	}
}
