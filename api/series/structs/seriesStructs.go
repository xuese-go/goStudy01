/**
系列管理
*/
package structs

import (
	"time"
)

type SeriesStructs struct {
	Uuid           string    `json:"uuid" gorm:"primary_key"` //gorm:"primary_key"声明为主键
	CreateTime     time.Time `json:"createTime" form:"createTime"`
	LastUpdateTime time.Time `json:"lastUpdateTime" form:"lastUpdateTime"` //最后修改时间
	Name           string    `json:"name" form:"name"`
}

//更改表名称
func (SeriesStructs) TableName() string {
	return "series_table"
}
