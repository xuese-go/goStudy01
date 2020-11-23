package structs

import (
	"github.com/xuese-go/goStudy01/db"
	"time"
)

/*
用户
*/
type UserStruct struct {
	Uuid           string    `json:"uuid" gorm:"primary_key"`              //gorm:"primary_key"声明为主键
	Account        string    `json:"account" gorm:"unique" form:"account"` //gorm:"unique"唯一
	Password       string    `json:"password" form:"password"`
	Role           uint      `json:"role" form:"role"`   //-1 管理员 1普通
	State          uint      `json:"state" form:"state"` //-1 停用 1正常
	CreateTime     time.Time `json:"createTime" form:"createTime"`
	LastUpdateTime time.Time `json:"lastUpdateTime" form:"lastUpdateTime"` //最后修改时间
}

func init() {
	//自动迁移表  默认不支持事务，所以加参数
	db.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&UserStruct{})
}

//更改表名称
func (UserStruct) TableName() string {
	return "user_table"
}
