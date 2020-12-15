package service

import (
	uuid "github.com/satori/go.uuid"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/api/user/structs"
	"github.com/xuese-go/goStudy01/db"
	"github.com/xuese-go/goStudy01/log"
	"github.com/xuese-go/goStudy01/util/md5"
	util "github.com/xuese-go/goStudy01/util/page"
	"time"
)

/**
新增用户
*/
func Save(user structs.UserStruct) resp.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()
	//查询是否有重复账号
	var u structs.UserStruct
	d := dba.First(&u, "account = ?", user.Account)
	if u.Account != "" {
		return resp.ResponeStruct{Success: false, Msg: "账号重复"}
	}

	//填充其它数据
	uid := uuid.NewV4().String()
	user.Uuid = uid
	user.CreateTime = time.Now()
	user.Role = 1
	user.State = 1
	//密码加密处理
	user.Password = md5.Enc(user.Password, user.Account)

	//新增数据
	t := tx.Create(user)
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
	var u structs.UserStruct
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
func Update(user structs.UserStruct) resp.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()
	var u structs.UserStruct
	if err := dba.First(&u, "uuid = ?", user.Uuid).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "查询错误"}
	}
	if user.Role != 0 {
		u.Role = user.Role
	}
	if user.State != 0 {
		u.State = user.State
	}
	if user.Password != "" {
		//重置密码
		if user.Password == "rest" {
			//密码加密处理
			u.Password = md5.Enc("111111", u.Account)
		} else {
			u.Password = md5.Enc(user.Password, u.Account)
		}
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
更改头像
*/
func UpdateImg(src string, uuid string) resp.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()
	var u structs.UserStruct
	if err := dba.First(&u, "uuid = ?", uuid).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "查询错误"}
	}
	u.Portrait = src
	u.LastUpdateTime = time.Now()
	t := tx.Save(&u)
	if t.Error != nil {
		t.Rollback()
		log.SugarLogger.Errorf(t.Error.Error())
		return resp.ResponeStruct{Success: false, Msg: "失败"}
	}
	t.Commit()
	return resp.ResponeStruct{Success: true, Msg: "成功", Data: src}
}

/**
根据id查询
*/
func One(uuid string) resp.ResponeStruct {
	dba := db.Db
	var u structs.UserStruct
	if err := dba.First(&u, "uuid = ?", uuid).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "查询错误"}
	}
	u.Password = ""
	return resp.ResponeStruct{Success: true, Msg: "操作成功", Data: u}
}

/**
根据id 获取是否停用
*/
func IsState(uuid string) resp.ResponeStruct {
	dba := db.Db
	var u structs.UserStruct
	if err := dba.First(&u, "uuid = ?", uuid).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "查询错误"}
	}
	u.Password = ""
	return resp.ResponeStruct{Success: u.State == 2}
}

/**
根据id 获取是否管理员
*/
func IsRole(uuid string) resp.ResponeStruct {
	dba := db.Db
	var u structs.UserStruct
	if err := dba.First(&u, "uuid = ?", uuid).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "查询错误"}
	}
	u.Password = ""
	return resp.ResponeStruct{Success: u.Role == 2}
}

/**
根据账号查询
*/
func ByAccount(account string) resp.ResponeStruct {
	dba := db.Db
	var u structs.UserStruct
	if err := dba.First(&u, "account = ?", account).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "账号或密码错误"}
	}
	return resp.ResponeStruct{Success: true, Msg: "操作成功", Data: u}
}

/**
分页查询
*/
func Page(pageNum int, pageSize int, user structs.UserStruct) resp.ResponeStruct {
	//为了不影响后边的操作  所以需要使用新的变量
	dba := db.Db
	us := make([]structs.UserStruct, 0)

	//查询条件
	if user.Account != "" {
		dba = dba.Where("account like ?", "%"+user.Account+"%")
	}
	if user.Role != 0 {
		dba = dba.Where("role = ?", user.Role)
	}
	if user.State != 0 {
		dba = dba.Where("state = ?", user.State)
	}
	//总记录数
	var count int
	dba = dba.Find(&us).Count(&count)
	if dba.Error != nil {
		log.SugarLogger.Errorf(dba.Error.Error())
		return resp.ResponeStruct{Success: false, Msg: "操作失败"}
	}

	//分页信息
	if pageNum > 0 && pageSize > 0 {
		dba = dba.Order("account")
		dba = dba.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}
	//查询
	if err := dba.Table("user_table").Select([]string{"uuid", "account", "role", "state", "portrait"}).Scan(&us).Error; err != nil {
		//if err := dba.Find(&us).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	return resp.ResponeStruct{Success: true, Data: us, Page: util.PageUtil(count, pageSize, pageNum)}
}
