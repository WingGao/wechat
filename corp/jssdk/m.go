package jssdk

import (
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/jssdk"
)

func NewDefaultTicketServer(clt core.IClient) (srv *jssdk.DefaultTicketServer) {
	srv = jssdk.NewDefaultTicketServer(clt)
	srv.SetTicketApi("https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=")
	return
}
