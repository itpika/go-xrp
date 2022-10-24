package rpc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/itpika/go-xrp/crypto"
	"github.com/itpika/go-xrp/data"
	"github.com/itpika/go-xrp/tools/http"
)

const (
	default_currency = "XRP"
)

// https://developers.ripple.com/sign.html
func (c *Client) Sign(from, to string, tag uint32, currency, value, fee, privateKey string, accountSequence, lastLedgerSequence uint32) (string, string, error) {
	fromAccount, _ := data.NewAccountFromAddress(from)
	toAccount, _ := data.NewAccountFromAddress(to)
	a := value
	if currency != "" {
		a += "/" + currency
	}
	amount, _ := data.NewAmount(a)
	feeVal, _ := data.NewValue(fee, true)

	txnBase := data.TxBase{
		TransactionType:    data.PAYMENT,
		Account:            *fromAccount,
		Sequence:           accountSequence,
		Fee:                *feeVal,
		LastLedgerSequence: &lastLedgerSequence,
	}
	var payment *data.Payment
	if tag > 0 {
		payment = &data.Payment{
			TxBase:         txnBase,
			Destination:    *toAccount,
			Amount:         *amount,
			DestinationTag: &tag,
		}
	} else {
		payment = &data.Payment{
			TxBase:      txnBase,
			Destination: *toAccount,
			Amount:      *amount,
		}
	}

	txid, txBlob, err := c.signOffline(payment, privateKey)
	if err != nil {
		return "", "", err
	}
	return txid, txBlob, nil
}

func (c *Client) signOffline(payment *data.Payment, privateKey string) (string, string, error) {
	pri, _ := hex.DecodeString(privateKey)
	key := crypto.LoadECDSKey(pri)

	err := data.Sign(payment, key, nil)
	if err != nil {
		return "", "", err
	}
	return c.makeTxBlob(payment)
}

// MakeTxblob
func (c *Client) makeTxBlob(payment *data.Payment) (string, string, error) {
	//fmt.Println("sign pub key: ", payment.SigningPubKey.String())
	h, raw, err := data.Raw(data.Transaction(payment))
	if err != nil {
		return "", "", err
	}
	txBlob := fmt.Sprintf("%X", raw)
	txid := h.String()
	return txid, txBlob, nil
}

func (c *Client) Submit(txBlob string) (*SubmitResult, error) {
	res := &SubmitResp{}
	params := `{"method": "submit", "params": [{"tx_blob": "` + txBlob + `"}]}`
	resp, err := http.HttpPost(c.rpcJsonURL, []byte(params))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, res)
	return res.Result, err
}

// https://developers.ripple.com/tx.html
func (c *Client) TX(hash string) (*TxResult, error) {
	params := `{"method":"tx", "params": [{"transaction":"` + hash + `"}]}`
	resp, err := http.HttpPost(c.rpcJsonURL, []byte(params))
	if err != nil {
		return nil, err
	}
	res := &TxResp{}
	err = json.Unmarshal(resp, res)
	return res.Result, nil
}
