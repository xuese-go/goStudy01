/**
酒管理
*/
package structs

import (
	"time"
)

type AlcoholStructs struct {
	Uuid           string    `json:"uuid" gorm:"primary_key"` //gorm:"primary_key"声明为主键
	CreateTime     time.Time `json:"createTime" form:"createTime"`
	LastUpdateTime time.Time `json:"lastUpdateTime" form:"lastUpdateTime"` //最后修改时间

	Name string `json:"name" form:"name"` //名称

	BrandId string `json:"brandId" form:"brandId"` //品牌

	SeriesId string `json:"seriesId" form:"seriesId"` //系列

	Concentration float32 `json:"concentration" form:"concentration"` //酒精浓度

}

//更改表名称
func (AlcoholStructs) TableName() string {
	return "alcohol_table"
}
