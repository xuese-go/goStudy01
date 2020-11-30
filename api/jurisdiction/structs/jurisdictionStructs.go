package structs

import (
	"time"
)

/*
权限
*/
type JurisdictionStruct struct {
	Uuid           string    `json:"uuid" gorm:"primary_key"`               //gorm:"primary_key"声明为主键
	JurName        string    `json:"jurName" gorm:"jurName" form:"jurName"` //gorm:"unique"唯一
	JurFlag        string    `json:"jurFlag" gorm:"jurFlag" form:"jurFlag"` //gorm:"unique"唯一
	CreateTime     time.Time `json:"createTime" form:"createTime"`
	LastUpdateTime time.Time `json:"lastUpdateTime" form:"lastUpdateTime"` //最后修改时间
}

//更改表名称
func (JurisdictionStruct) TableName() string {
	return "jurisdiction_table"
}
