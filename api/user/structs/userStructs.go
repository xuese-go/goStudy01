package structs

import (
	"time"
)

/*
用户
*/
type UserStruct struct {
	Uuid           string    `json:"uuid" gorm:"primary_key"`              //gorm:"primary_key"声明为主键
	Account        string    `json:"account" gorm:"unique" form:"account"` //gorm:"unique"唯一
	Password       string    `json:"password" form:"password"`
	Role           int       `json:"role" form:"role"`   //2 管理员 1普通
	State          int       `json:"state" form:"state"` //2 停用 1正常
	CreateTime     time.Time `json:"createTime" form:"createTime"`
	LastUpdateTime time.Time `json:"lastUpdateTime" form:"lastUpdateTime"` //最后修改时间
}

//更改表名称
func (UserStruct) TableName() string {
	return "user_table"
}
