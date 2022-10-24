package rpc

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/cosmos/btcutil"
	"github.com/itpika/go-xrp/crypto"
	"github.com/mr-tron/base58"
)

const (
	TestNet = "https://s.altnet.rippletest.net:51234"
	DevNet  = "https://s.devnet.rippletest.net:51234"
	MainNet = "https://s1.ripple.com:51234/"
)

var (
	client = NewClient(TestNet)
	//client = NewClient("http://47.75.70.201:9003", "http://47.75.70.201:9003")
	//client = NewClient("https://data.ripple.com")
)

func TestAddress(t *testing.T) {
	priv := "71d937e941203f76e5da4ae1fcba049bf2a7c80a002d8d9c60b69a21919390f9"
	priBt, err := hex.DecodeString(priv)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(crypto.LoadECDSKey(priBt).EncodeAddressString(true))
	fmt.Println(hex.EncodeToString(crypto.LoadECDSKey(priBt).Public(nil)))
}

func TestPubToAddr(t *testing.T) {

	// pri: 71d937e941203f76e5da4ae1fcba049bf2a7c80a002d8d9c60b69a21919390f9
	// pub: 02a70106c3935d7e5a607f523dd3dc46be7a061a0e424f91b3ccc36b37973b0325
	bt, err := hex.DecodeString("02a70106c3935d7e5a607f523dd3dc46be7a061a0e424f91b3ccc36b37973b0325")
	if err != nil {
		t.Fatal(err)
	}
	hash160 := btcutil.Hash160(bt)
	b := make([]byte, 0, 1+len(hash160)+4)
	b = append(b, 0)
	b = append(b, hash160...)

	h := sha256.Sum256(b)
	h2 := sha256.Sum256(h[:])
	checkSum := h2[:4]
	b = append(b, checkSum...)

	// 0067d91565dff3aed784edc56abe741c5ca80000b7c542ba8e
	fmt.Println(hex.EncodeToString(b))
	fmt.Println(base58.EncodeAlphabet(b, base58.NewAlphabet(crypto.ALPHABET)))

}

func TestGetAccountInfo(t *testing.T) {
	address := "r97g8NCTZpPt5vacawMvbeoeh3fYSSuJLq"
	res, err := client.GetAccountInfo(address)
	if err != nil {
		t.Error("err: ", err)
	}
	fmt.Printf("res: %+v\n", res.AccountData)
}

func TestGenAddress(t *testing.T) {
	pri, pub, addr, err := client.GenAddress()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("pri: ", pri)
	fmt.Println("pub: ", pub)
	fmt.Println("addr: ", addr)
}

//pri:  d3c34cc4553591860f14fb64dd9562210f57b2e12970e752e54402fc7dd2844f
//pub:  0373330fcc500d6e7b1ce775ac9ca2cfa13b805befcd5b17d5108d1246d1bb6337
//addr:  raMjJMN8LDogUx6BckDV7LojR8XAXZDm91

//pri:  c57891a6f2212dd312a12cb9323e69b6ad8a0faaf8435ca533876a7c12b80ae8
//pub:  033b849512a08922e93c74d43c13c6b1c2dc8591bf787b4f36e81d511bea587dd6
//addr:  rG5AB117rJ7e2MZGKE4XfaVK5BdyHBxcSm
