package rpc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/yancaitech/go-xrp/crypto"
	"github.com/yancaitech/go-xrp/tools/http"
)

// GenAddress 生成账户地址
func (c *Client) GenAddress() (string, string, string, error) {
	key, err := crypto.GenEcdsaKey()
	if err != nil {
		return "", "", "", err
	}
	var seq0 uint32
	address, err := crypto.AccountId(key, &seq0)
	pri := hex.EncodeToString(key.Private(&seq0))
	pub := hex.EncodeToString(key.Public(&seq0))
	return pri, pub, address.String(), err
}

// GetAccountBalances 获取账户余额
func (c *Client) GetAccountBalances(address string, queryParams map[string]string) (*AccountBalancesStruct, error) {
	balance := &AccountBalancesStruct{}
	if address == "" {
		return balance, fmt.Errorf("address is empty")
	}
	host := "/v2/accounts/" + address + "/balances"
	values := make(url.Values)
	if queryParams != nil {
		for key, val := range queryParams {
			values.Add(key, val)
		}
	}
	queryUrl := c.apiURL + host + values.Encode()
	resp, err := http.HttpGet(queryUrl)
	if err != nil {
		return balance, err
	}
	err = json.Unmarshal(resp, balance)
	if err != nil {
		return balance, err
	}
	if balance.Result != "success" {
		return balance, fmt.Errorf(balance.Message)
	}
	return balance, nil
}

func (c *Client) GetAccountInfo(address string) (*AccountInfoResult, error) {
	params := map[string]interface{}{
		"method": "account_info",
		"params": []map[string]string{
			{
				"account": address,
				"ledgder": "validated",
			},
		},
	}
	str, _ := json.Marshal(params)
	resp, err := http.HttpPost(c.rpcJsonURL, str)
	if err != nil {
		return nil, err
	}
	res := &AccountInfoResp{}
	err = json.Unmarshal(resp, res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

// GetAccountTransactions https://developers.ripple.com/data-api.html#get-account-transaction-history
func (c *Client) GetAccountTransactions(address string, params map[string]string) (*AccountTransactionResp, error) {
	if address == "" {
		return nil, nil
	}
	uri := c.apiURL + "/v2/accounts/" + address + "/transactions"
	query := make(url.Values)
	if params != nil {
		for k, v := range params {
			query.Add(k, v)
		}
	}
	if len(query) > 0 {
		uri = uri + "?" + query.Encode()
	}
	resp, err := http.HttpGet(uri)
	if err != nil {
		return nil, err
	}
	res := &AccountTransactionResp{}
	err = json.Unmarshal(resp, res)
	return res, err
}
