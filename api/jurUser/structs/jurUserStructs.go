package structs

import (
	"time"
)

/*
用户权限
*/
type JurUserStruct struct {
	Uuid           string    `json:"uuid" gorm:"primary_key"`
	JurId          string    `json:"jurId" gorm:"jurId" form:"jurId"`
	UserId         string    `json:"userId" gorm:"userId" form:"userId"`
	CreateTime     time.Time `json:"createTime" form:"createTime"`
	LastUpdateTime time.Time `json:"lastUpdateTime" form:"lastUpdateTime"`
}

//更改表名称
func (JurUserStruct) TableName() string {
	return "jur_user_table"
}
