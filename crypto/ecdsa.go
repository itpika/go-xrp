package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/cosmos/btcutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

var (
	order = btcec.S256().N
	zero  = big.NewInt(0)
	one   = big.NewInt(1)
)

type EcdsaKey struct {
	*btcec.PrivateKey
}

func newKey(seed []byte) *btcec.PrivateKey {
	inc := big.NewInt(0).SetBytes(seed)
	inc.Lsh(inc, 32)
	for key := big.NewInt(0); ; inc.Add(inc, one) {
		key.SetBytes(Sha512Half(inc.Bytes()))
		if key.Cmp(zero) > 0 && key.Cmp(order) < 0 {
			privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), key.Bytes())
			return privKey
		}
	}
}

// If seed is nil, generate a random one
func NewECDSAKey(seed []byte) (*EcdsaKey, error) {
	if seed == nil {
		seed = make([]byte, 16)
		if _, err := rand.Read(seed); err != nil {
			return nil, err
		}
	}
	return &EcdsaKey{newKey(seed)}, nil
}

// GenEcdsaKey 生成 secp256k1 key
func GenEcdsaKey() (*EcdsaKey, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	k := (*btcec.PrivateKey)(key)
	return &EcdsaKey{k}, nil
}
func LoadECDSKey(privateKeyBytes []byte) *EcdsaKey {

	pri, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBytes)
	k := &EcdsaKey{pri}
	return k
}

func (k *EcdsaKey) generateKey(sequence uint32) *btcec.PrivateKey {
	seed := make([]byte, btcec.PubKeyBytesLenCompressed+4)
	copy(seed, k.PubKey().SerializeCompressed())
	binary.BigEndian.PutUint32(seed[btcec.PubKeyBytesLenCompressed:], sequence)
	key := newKey(seed)
	key.D.Add(key.D, k.D).Mod(key.D, order)
	key.X, key.Y = key.ScalarBaseMult(key.D.Bytes())
	return key
}

func (k *EcdsaKey) Id(sequence *uint32) []byte {
	if sequence == nil {
		return Sha256RipeMD160(k.PubKey().SerializeCompressed())
	}
	return Sha256RipeMD160(k.Public(sequence))
}

func (k *EcdsaKey) Private(sequence *uint32) []byte {
	if sequence == nil {
		return k.D.Bytes()
	}
	return k.generateKey(*sequence).D.Bytes()
}

func (k *EcdsaKey) Public(sequence *uint32) []byte {
	if sequence == nil {
		return k.PubKey().SerializeCompressed()
	}
	return k.generateKey(*sequence).PubKey().SerializeCompressed()
}

func (k *EcdsaKey) EncodeAddressString(compressed bool) string {
	hash160 := btcutil.Hash160(k.PubKey().SerializeCompressed())
	b := make([]byte, 0, 1+len(hash160)+4)
	b = append(b, 0)
	b = append(b, hash160...)

	h := sha256.Sum256(b)
	h2 := sha256.Sum256(h[:])
	checkSum := h2[:4]
	b = append(b, checkSum...)

	return base58.EncodeAlphabet(b, base58.NewAlphabet(ALPHABET))
}
