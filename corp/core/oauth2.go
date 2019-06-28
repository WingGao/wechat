package core

import (
	mpcore "github.com/chanxuehong/wechat/mp/core"
)

type UserAuthInfo struct {
	UserId   string `json:"UserId"` // 成员UserID。若需要获得用户详情信息，可调用通讯录接口：读取成员
	OpenId   string `json:"OpenId"` // 非企业成员的标识，对当前企业唯一
	DeviceId string `json:"DeviceId"`
}

// 该接口用于根据code获取成员信息
// https://work.weixin.qq.com/api/doc#90000/90135/91023
func (c *Client) GetUserAuthInfo(code string) (*UserAuthInfo, error) {
	var result struct {
		mpcore.Error
		UserAuthInfo
	}
	err := c.GetJSON("https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?code="+code+"&access_token=", &result)
	if err != nil {
		return nil, err
	}
	return &result.UserAuthInfo, err
}
