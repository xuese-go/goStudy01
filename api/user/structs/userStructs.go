package structs

import "time"

/*
用户
*/
type UserStruct struct {
	Uuid     string `json:"uuid"`
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
	//-1 管理员 0普通
	Role int `json:"role"`
	//-1 停用 0正常
	State      int       `json:"state"`
	CreateTime time.Time `json:"createTime"`
	//最后修改时间
	LastUpdateTime time.Time `json:"lastUpdateTime"`
}
