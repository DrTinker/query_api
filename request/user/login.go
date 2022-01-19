package user

import (
	"query_api/models"
)

// 实现登录参数校验
func LoginDataCheck(user models.User) bool {
	u := user
	if (u.User_id >= 0 && u.User_pwd != "") {
		return false
	}
	return true
}