package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"

	"crypto/ed25519"

	"github.com/cosmos/btcutil"
	"github.com/yancaitech/go-xrp/crypto"
)

type ed25519key struct {
	priv [ed25519.PrivateKeySize]byte
}

func checkSequenceIsNil(seq *uint32) {
	if seq != nil {
		panic("Ed25519 keys do not support account families")
	}
}

func (e *ed25519key) Id(seq *uint32) []byte {
	checkSequenceIsNil(seq)
	return Sha256RipeMD160(e.Public(seq))
}

func (e *ed25519key) Public(seq *uint32) []byte {
	checkSequenceIsNil(seq)
	return append([]byte{0xED}, e.priv[32:]...)
}

func (e *ed25519key) Private(seq *uint32) []byte {
	checkSequenceIsNil(seq)
	return e.priv[:]
}

func NewEd25519Key(seed []byte) (*ed25519key, error) {
	r := rand.Reader
	if seed != nil {
		r = bytes.NewReader(Sha512Half(seed))
	}
	_, priv, err := ed25519.GenerateKey(r)
	if err != nil {
		return nil, err
	}
	var k ed25519key
	copy(k.priv[:], priv)
	return &k, nil
}

func (k *ed25519key) EncodeAddressString() string {

	hash160 := btcutil.Hash160(k.Public(nil))
	b := make([]byte, 0, 1+len(hash160)+4)
	b = append(b, 0)
	b = append(b, hash160...)

	h := sha256.Sum256(b)
	h2 := sha256.Sum256(h[:])
	checkSum := h2[:4]
	b = append(b, checkSum...)

	return crypto.Base58Encode(b, crypto.ALPHABET)
}
