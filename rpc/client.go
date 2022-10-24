package rpc

type Client struct {
	rpcJsonURL string
}

func NewClient(rpcJsonURL string) *Client {
	return &Client{
		rpcJsonURL: rpcJsonURL,
	}
}
