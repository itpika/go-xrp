package rpc

import (
	"fmt"
	"testing"
)

func TestGetServerInfo(t *testing.T) {
	res, err := client.GetServerInfo()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.State.ValidatedLedger)
}

func TestGetServerFee(t *testing.T) {
	res, err := client.GetServerFee()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.Drops)
}
