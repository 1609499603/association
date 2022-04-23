package models

// Association 社团表
type Association struct {
	//id
	Id int64 `json:"id"`
	//社团名称
	Name string `json:"name"`
	//简介
	Introduction string `json:"introduction"`
	//招生宣言
	EnrollmentDeclaration string `json:"enrollment_declaration"`

	//逻辑删除，记录时间
	Record
}
