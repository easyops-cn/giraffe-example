package logic_cmdb_service

import (
	"github.com/easyops-cn/giraffe-micro"

	"github.com/easyops-cn/giraffe-example/protorepo_logic_cmdb_service/instances"
)

type Client struct {
	Instances instances.Client
}

func NewClient(c giraffe.Client) *Client {
	client := Client{}
	client.Instances = instances.NewClient(c)
	return &client
}
