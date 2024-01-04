package client

import "fmt"

type Client struct {
}

func (*Client) Noop() {
	fmt.Println("Hey!")
}
