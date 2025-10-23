package Local

import (
	"RSAandAES"
	"encoding/json"
	"math/big"
)

type SignedTransaction struct {
	Tx  *Transaction
	Sgn *big.Int
}

func (c *Client) SignTransaction(tx *Transaction) (*SignedTransaction, error) {
	txInBytes, err := json.Marshal(tx)
	if err != nil {
		return new(SignedTransaction), err
	}
	sk := StringToSecretKey(c.SecretKey)
	sgn := RSAandAES.GenerateCertificate(txInBytes, sk.N, sk.D)

	sgn_tx := new(SignedTransaction)
	sgn_tx.Tx = tx
	sgn_tx.Sgn = sgn

	return sgn_tx, nil
}

func VerifySignedTransaction(sgnTx *SignedTransaction) (bool, error) {
	tx := sgnTx.Tx
	sgn := sgnTx.Sgn

	txInBytes, err := json.Marshal(tx)
	if err != nil {
		return false, err
	}

	pkString := tx.From
	pk := StringToPublicKey(pkString)
	n := pk.N
	e := pk.E

	return RSAandAES.VerifyCertificate(txInBytes, n, e, sgn), nil
}
