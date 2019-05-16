package pool

import (
	"github.com/seefan/gossdb/client"
)

//Client pooled client
type Client struct {
	client.Client
	//连接是否正在使用，防止重复关闭
	used bool
	//pool 中的位置
	index int
	//连接池块
	pool *Pool
	//连接池
	over *Connectors
}

//Close put the client to Connectors
func (c *Client) Close() {
	if c.Error == nil {
		if c.used {
			c.over.closeClient(c)
		}
	} else {
		c.over.poolTemp.Put(c)
	}
}