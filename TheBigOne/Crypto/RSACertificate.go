package RSAandAES

import (
	"crypto/sha256"
	"math/big"
)

func GenerateCertificate(msg []byte, modulus *big.Int, secretKey *big.Int) *big.Int {
	hashedMsg := sha256.Sum256(msg)
	var hashedMsgInt *big.Int = new(big.Int)
	hashedMsgInt.SetBytes(hashedMsg[:])
	var certificate *big.Int = Decrypt(hashedMsgInt, modulus, secretKey)
	return certificate
}

func VerifyCertificate(msg []byte, modulus *big.Int, publicKey *big.Int, certificate *big.Int) bool {
	hashedMsg := sha256.Sum256(msg)
	var hashedMsgInt *big.Int = new(big.Int)
	hashedMsgInt.SetBytes(hashedMsg[:])

	decipheredCertificate := Encrypt(certificate, modulus, publicKey)

	return hashedMsgInt.Cmp(decipheredCertificate) == 0
}
