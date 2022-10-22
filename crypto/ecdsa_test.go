package crypto

import (
	"fmt"
	"testing"
)

func TestGenKey(t *testing.T) {
	// secret := "snh1zUj8AKjdLPDRapFGpJeaBRDHm"
	secret := "sEdVZ4zKBYahvGoDGX2RdxoKNe1PVh8"
	seed, err := NewRippleHash(secret)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%x\n", seed.Payload())
	key, err := NewECDSAKey(seed.Payload())
	if err != nil {
		t.Fatal(err)
	}
	var sequenceZero uint32
	h, err := AccountPrivateKey(key, &sequenceZero)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("hash: ", h.String())
	h, _ = AccountPublicKey(key, &sequenceZero)
	fmt.Println("hash: ", h.String())
	fmt.Printf("pri: %x\n", key.Private(&sequenceZero))
	fmt.Printf("pub: %x\n", key.Public(&sequenceZero))

}
