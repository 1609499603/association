package models

import "time"

// Record 相关操作
type Record struct {
	//逻辑删除 0未删除，1已删除
	IsDeleted int `json:"is_deleted"`
	//创建时间
	CreateTime time.Time `json:"create_time"`
	//更新时间
	UpdateTime time.Time `json:"update_time"`
}
