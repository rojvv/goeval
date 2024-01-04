package client

import "fmt"

type Client struct {
}

func (*Client) Noop() {
	fmt.Println("Hey!")
}

func NewClient() *Client {
	return &Client{}
}
