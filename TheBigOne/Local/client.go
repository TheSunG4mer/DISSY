package Local

import (
	"RSAandAES"
	"math/big"
	"strings"
)

type Client struct {
	PublicKey string
	SecretKey string
}

type PublicKey struct {
	N *big.Int
	E *big.Int
}

type SecretKey struct {
	N *big.Int
	D *big.Int
}

func PublicKeyToString(pk *PublicKey) string {
	outString := pk.N.String() + "," + pk.E.String()
	return outString
}

func StringToPublicKey(s string) *PublicKey {
	sliceOfString := strings.Split(s, `,`)
	N := big.NewInt(0)
	N.SetString(sliceOfString[0], 10)
	E := big.NewInt(0)
	E.SetString(sliceOfString[1], 10)
	return &PublicKey{N, E}
}

func SecretKeyToString(sk *SecretKey) string {
	outString := sk.N.String() + "," + sk.D.String()
	return outString
}

func StringToSecretKey(s string) *SecretKey {
	sliceOfString := strings.Split(s, `,`)
	N := big.NewInt(0)
	N.SetString(sliceOfString[0], 10)
	D := big.NewInt(0)
	D.SetString(sliceOfString[1], 10)
	return &SecretKey{N, D}
}

func MakeClient(k int) (*Client, error) {
	n, e, d, err := RSAandAES.KeyGen(k)
	if err != nil {
		return new(Client), err
	}
	pk := PublicKey{n, e}
	sk := SecretKey{n, d}

	c := Client{PublicKeyToString(&pk), SecretKeyToString(&sk)}
	return &c, nil

}
