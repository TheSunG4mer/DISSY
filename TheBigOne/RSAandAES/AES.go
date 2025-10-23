package RSAandAES

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"os"
)

func AESKeyGen(k int) ([]byte, error) {
	key := make([]byte, k)
	rand.Read(key)
	return key, nil
}

func EncryptToFile(msg []byte, fileName string, key []byte) ([]byte, error) {
	for len(msg)%16 != 0 {
		msg = append(msg, 0)
	}
	message_length := len(msg)

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	output := []byte{}

	iv := make([]byte, 16)
	rand.Read(iv)
	output = append(output, iv...)

	for i := 0; 16*i < message_length; i++ {
		toXOR := make([]byte, 16)
		iv[15] += 1
		cipher.Encrypt(toXOR, iv)
		for j := range 16 {
			toXOR[j] = toXOR[j] ^ msg[16*i+j]
		}
		output = append(output, toXOR...)
	}

	fo, err := os.Create(fileName)
	if err != nil {
		return []byte{}, err
	}
	defer fo.Close()

	_, err = fo.Write(output)
	if err != nil {
		return []byte{}, err
	}
	fmt.Println("File opened")

	return output, nil
}

func DecryptFromFile(cipherText []byte, key []byte) ([]byte, error) {
	// msgFromFile := []byte{}
	// fi, err := os.Open(fileName)
	// if err != nil {
	// 	return []byte{}, err
	// }

	// c := make([]byte, 16)
	// for {

	// 	n, err := fi.Read(c)
	// 	if err != nil {
	// 		return []byte{}, err
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}
	// 	msgFromFile = append(msgFromFile, c...)
	// }
	// fi.Close()

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	iv := cipherText[:16]
	output := []byte{}
	fmt.Println("Decrypting")

	for i := 0; 16*i+16 < len(cipherText); i++ {
		toXOR := make([]byte, 16)
		iv[15] += 1
		cipher.Encrypt(toXOR, iv)
		for j := range 16 {
			toXOR[j] = toXOR[j] ^ cipherText[16*i+16+j]
		}
		output = append(output, toXOR...)
	}

	return output, nil
}
