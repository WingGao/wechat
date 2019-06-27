package core

import (
	mpcore "github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/util"
	"net/http"
)

type Client struct {
	*mpcore.Client
	HttpClient *http.Client
}

// NewClient 创建一个新的 Client.
//  如果 clt == nil 则默认用 util.DefaultHttpClient
func NewClient(srv mpcore.AccessTokenServer, clt *http.Client) *Client {
	if srv == nil {
		panic("nil AccessTokenServer")
	}
	if clt == nil {
		clt = util.DefaultHttpClient
	}
	return &Client{
		Client:     mpcore.NewClient(srv, clt),
		HttpClient: clt,
	}
}
