package rpc

type Client struct {
	rpcJsonURL string
}

func NewClient(rpcJsonURL, apiURL string) *Client {
	return &Client{
		rpcJsonURL: rpcJsonURL,
	}
}
