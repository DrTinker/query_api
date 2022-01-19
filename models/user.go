package models

type User struct {
	User_name  string `json:"user_name" form:"user_name"`
	User_id    int32  `json:"user_id" form:"user_id"`
	User_pwd   string `json:"user_pwd" form:"user_pwd"`
	User_tag   string `json:"user_tag" form:"user_tag"` // 兴趣标签的json字符串
	Log        string `json:"log" form:"log"`           // 用户备注
	Pass       bool   `json:"pass" form:"pass"`
	Woekspace  string `json:"woekspace" form:"woekspace"` // 工作空间包含问卷ID的json字符串
	History    string `json:"history" form:"history"`     // 历史问卷包含问卷ID的json字符串
	Subscribe  string `json:"subscribe" form:"subscribe"` // 收藏夹包含问卷ID的json字符串
	User_phone int    `json:"user_phone" form:"user_phone"`
	User_email string `json:"user_email" form:"user_email"`
}

type Login struct {
	User_id    int32  `json:"user_id" form:"user_id"`
	User_pwd   string `json:"user_pwd" form:"user_pwd"`
}
