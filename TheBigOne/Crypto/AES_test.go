package RSAandAES

import (
	"testing"
)

func TestAES(t *testing.T) {
	k, err := AESKeyGen(16)
	if err != nil {
		t.Errorf("Error occured during Key generation")
	}

	msg := []byte("Secret message, secret message, secret message")

	fileName := "C:/Users/au649790/OneDrive - Aarhus universitet/Desktop/DISSY/TheBigOne/Crypto/cipher.txt"

	cipher, err := EncryptToFile(msg, fileName, k)
	if err != nil {
		t.Errorf("Error occured during encryption: %v", err)
	}

	new_msg, err := DecryptFromFile(cipher, k)
	if err != nil {
		t.Errorf("Error occured during decryption %v\n", err)
	}

	for j := range len(msg) {
		if msg[j] != new_msg[j] {
			t.Error("Got the wrong message out")
		}
	}
}
