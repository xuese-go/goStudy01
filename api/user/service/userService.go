package service

import (
	uuid "github.com/satori/go.uuid"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/api/user/structs"
	"github.com/xuese-go/goStudy01/db"
	"time"
)

/**
新增用户
*/
func Save(user structs.UserStruct) resp.ResponeStruct {

	//查询是否有重复账号
	var u structs.UserStruct
	d := db.Db.First(&u, "account = ?", user.Account)
	if u.Account != "" {
		return resp.ResponeStruct{Success: false, Msg: "账号重复"}
	}

	//填充其它数据
	uid := uuid.NewV4().String()
	user.Uuid = uid
	user.CreateTime = time.Now()
	user.Role = 0
	user.State = 0

	//新增数据
	d = db.Db.Create(user)
	if d.Error != nil {
		d.Rollback()
		return resp.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	return resp.ResponeStruct{Success: true, Msg: "操作成功"}
}

/**
分页查询
*/
func Page(pageNum int, pageSize int, user structs.UserStruct) resp.ResponeStruct {
	//us := make([]structs.UserStruct, 0)
	//a := ""
	////b := make([]string, 0)
	////if user.Account != "" {
	////	a += " and account like ?"
	////	b = append(b, "%"+user.Account+"%")
	////}
	//err := db.Db.Select(&us, sql.PAGE+a+sql.LIMIT, (pageNum-1)*pageSize,pageSize)
	//if err != nil {
	//	fmt.Println(err)
	//	return resp.ResponeStruct{Success: false, Msg: "操作失败"}
	//}
	return resp.ResponeStruct{Success: true, Data: nil}
}
