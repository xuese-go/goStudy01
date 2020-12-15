package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	jurStruct "github.com/xuese-go/goStudy01/api/jurisdiction/structs"
	resp "github.com/xuese-go/goStudy01/api/respone/structs"
	"github.com/xuese-go/goStudy01/db"
	"github.com/xuese-go/goStudy01/log"
)

/**
修改
*/
func Update(id string, jurs []string) resp.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()
	us := make([]jurStruct.JurisdictionStruct, 0)
	if err := tx.Find(&us, "userId = ?", id).Delete(&us).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		tx.Rollback()
		return resp.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	tx.Commit()
	//return resp.ResponeStruct{Success: true, Msg: "操作成功"}
	if jurs == nil || len(jurs) <= 0 {
		return resp.ResponeStruct{Success: true, Msg: "操作成功"}
	} else {
		//填充其它数据
		//生成sql
		insertSql := ""
		for i := range jurs {
			insertSql += fmt.Sprintf("insert into `jur_user_table` (uuid,jurId,userId) values ('%s','%s','%s');", uuid.NewV4().String(), jurs[i], id)
		}
		if err := tx.Exec(insertSql).Error; err != nil {
			tx.Rollback()
			log.SugarLogger.Errorf(tx.Error.Error())
			return resp.ResponeStruct{Success: false, Msg: "操作失败"}
		} else {
			tx.Commit()
			return resp.ResponeStruct{Success: true, Msg: "操作成功"}
		}
	}

}

/**
根据用户主键查询
@return []jurStruct.JurisdictionStruct
*/
func FindByUserId(uuid string) resp.ResponeStruct {
	if uuid == "" {
		return resp.ResponeStruct{Success: false, Msg: "参数不能为空"}
	}
	//为了不影响后边的操作  所以需要使用新的变量
	dba := db.Db
	us := make([]jurStruct.JurisdictionStruct, 0)
	dba = dba.Where("jur_user_table.userId = ?", uuid)

	//查询
	dba = dba.Table("jur_user_table").Select([]string{"uuid", "jur_name", "jur_flag"})
	dba = dba.Joins("left join jurisdiction_table on jurisdiction_table.uuid = jur_user_table.jurId")
	if err := dba.Scan(&us).Error; err != nil {
		log.SugarLogger.Errorf(err.Error())
		return resp.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	return resp.ResponeStruct{Success: true, Data: us}
}
