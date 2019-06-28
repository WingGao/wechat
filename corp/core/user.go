package core

import (
	mpcore "github.com/chanxuehong/wechat/mp/core"
)

type UserInfo struct {
	UserId string `json:"userid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Gender string `json:"gender"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}

// 读取成员
// https://work.weixin.qq.com/api/doc#90000/90135/90196
func (c *Client) GetUser(userId string) (*UserInfo, error) {
	var result struct {
		mpcore.Error
		UserInfo
	}
	err := c.GetJSON("https://qyapi.weixin.qq.com/cgi-bin/user/get?userid="+userId+"&access_token=", &result)
	if err != nil {
		return nil, err
	}
	return &result.UserInfo, err
}
