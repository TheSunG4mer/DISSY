package RSAandAES

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestRSAKeyGen(t *testing.T) {
	n, _, _, err := KeyGen(2000)
	// fmt.Println("Testing")
	if err != nil {
		t.Errorf("KeyGen returned non-nil error")
	}

	// fmt.Printf("%p\n", n)

	if n.BitLen() < 3990 {
		t.Errorf("Key generated was too short: got %d bits, want at least 3990", n.BitLen())
	}
}

func TestEncryptDecrypt(t *testing.T) {
	for range 100 { // make 100 random tests
		n, e, d, err := KeyGen(100)
		if err != nil {
			t.Errorf("Error during Keygeneration")
		}
		msg, err := rand.Int(rand.Reader, big.NewInt(400000000))
		if err != nil {
			t.Errorf("Error during message generation")
		}
		ciffer := Encrypt(msg, n, e)
		new_msg := Decrypt(ciffer, n, d)
		if new_msg.Cmp(msg) != 0 {
			t.Errorf("Error during encoding/decoding")
		}
	}
}
