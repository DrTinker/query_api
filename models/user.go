package models

type User struct {
	User_name  string   `json:"user_name" form:"user_name"`
	User_id    int32    `json:"user_id" form:"user_id"`
	User_pwd   string   `json:"user_pwd" form:"user_pwd"`
	User_tag   []string `json:"user_tag" form:"user_tag"` // 兴趣标签的json字符串
	Log        string   `json:"log" form:"log"`           // 用户备注
	Pass       bool     `json:"pass" form:"pass"`
	User_phone int      `json:"user_phone" form:"user_phone"`
	User_email string   `json:"user_email" form:"user_email"`
}

type Login struct {
	User_id  int32  `json:"user_id" form:"user_id"`
	User_pwd string `json:"user_pwd" form:"user_pwd"`
}

type Token struct {
	UserID  int32  `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	Expire  int64  `json:"expire"`
}
