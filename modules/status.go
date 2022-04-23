package models

type Status struct {
	//权限id
	Id int64 `json:"id"`
	//权限名称
	Status string `json:"status"`
}
