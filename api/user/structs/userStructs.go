package structs

import (
	"github.com/xuese-go/goStudy01/db"
	"time"
)

/*
用户
*/
type UserStruct struct {
	Uuid           string    `json:"uuid" db:"uuid"`
	Account        string    `json:"account" form:"account" db:"account"`
	Password       string    `json:"password" form:"password" db:"password"`
	Role           uint      `json:"role" form:"role" db:"role"`    //-1 管理员 0普通
	State          uint      `json:"state" form:"state" db:"state"` //-1 停用 0正常
	CreateTime     time.Time `json:"createTime" form:"createTime" db:"createTime"`
	LastUpdateTime time.Time `json:"lastUpdateTime" form:"lastUpdateTime" db:"lastUpdateTime"` //最后修改时间
}

func init() {
	//自动迁移表
	db.Db.AutoMigrate(&UserStruct{})
}

//更改表名称
func (UserStruct) TableName() string {
	return "user_table"
}
