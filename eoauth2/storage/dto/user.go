package dto

import (
	"encoding/json"
	"fmt"
)

// User 用户信息
type User struct {
	// 用户uid
	Uid int64 `json:"uid"`
	// 用户昵称，中文名
	Nickname string `json:"nickname"`
	// 用户名，拼音
	Username string `json:"username"`
	// 头像
	Avatar string `json:"avatar"`
	// 邮箱
	Email string `json:"email"`
}

func (u *User) Marshal() (string, error) {
	if u == nil {
		return "", fmt.Errorf("user is nil")
	}

	bytes, err := json.Marshal(u)
	return string(bytes), err
}
