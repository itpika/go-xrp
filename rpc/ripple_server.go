package rpc

import (
	"encoding/json"

	"github.com/itpika/go-xrp/tools/http"
)

func (c *Client) GetServerInfo() (*ServerInfoResult, error) {
	resp, err := http.HttpPost(c.rpcJsonURL, []byte(`{"method":"server_state", "params": [{}]}`))
	if err != nil {
		return nil, err
	}
	res := &ServerInfoResp{}
	err = json.Unmarshal(resp, res)
	return res.Result, err
}

func (c *Client) GetServerFee() (*FeeResult, error) {
	resp, err := http.HttpPost(c.rpcJsonURL, []byte(`{"method":"fee", "params": [{}]}`))
	if err != nil {
		return nil, err
	}
	res := &ServerFee{}
	err = json.Unmarshal(resp, res)
	return res.Result, err
}
