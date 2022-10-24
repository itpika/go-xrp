package rpc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/itpika/go-xrp/crypto"
	"github.com/itpika/go-xrp/tools/http"
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

func (c *Client) GetAccountInfo(address string) (*AccountInfoResult, error) {
	params := map[string]interface{}{
		"method": "account_info",
		"params": []map[string]any{
			{
				"account":      address,
				"strict":       true,
				"ledger_index": "current",
				"queue":        true,
			},
		},
	}
	str, _ := json.Marshal(params)
	resp, err := http.HttpPost(c.rpcJsonURL, str)
	if err != nil {
		return nil, err
	}
	res := &AccountInfoResp{}
	fmt.Println(string(resp))
	err = json.Unmarshal(resp, res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}
