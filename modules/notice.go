package models

// Notice 公告表
type Notice struct {
	//公告id
	Id int64 `json:"id"`
	//公告标题
	Title string `json:"title"`
	//公告内容
	Content string `json:"content"`
	//是否为平台通知
	IsSystem int `json:"is_system"`
	//所属社团id
	AssociationId int64 `json:"association_id"`

	//逻辑删除，记录时间
	Record
}
