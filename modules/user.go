package models

// User 用户表
type User struct {
	//用户id
	Id int64 `json:"id"`
	//账号
	Username string `json:"username"`
	//密码
	Password string `json:"password"`
	//角色（老师或学生，0学生，1老师）
	Role int `json:"role"`
	//身份(1普通，2副社长，3社长，4指导老师，5超级管理)
	StatusId int `json:"status_id"`
	//逻辑删除，记录时间
	Record
}

// OnlineUser 用户线上数据
type OnlineUser struct {
	Id            int64  `json:"id"`            //用户id
	Username      string `json:"username"`      //用户名
	LoginTime     int64  `json:"loginTime"`     //登录时间
	LoginLocation string `json:"loginLocation"` // 归属地
	Browser       string `json:"browser"`       // 浏览器
	College       string `json:"College"`       // 学院
	Ip            string `json:"ip"`            //ip地址
	Token         string `json:"key"`           // token
}

type LoginRes struct {
	Username  string `json:"username"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type Context struct {
	Username string `json:"username"`
	Password string `json:"password"`
	StatusId int    `json:"status_id"`
}
