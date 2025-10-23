package Local_test

import (
	"Local"
	"testing"
)

func TestVerifySignedTransaction(t *testing.T) {
	c, err := Local.MakeClient(200)
	if err != nil {
		t.Error(err)
	}

	tx := Local.MakeTransaction("1234", c.PublicKey, "Bob", 10)

	sgn_tx, err := c.SignTransaction(tx)
	if err != nil {
		t.Error(err)
	}

	ver, err := Local.VerifySignedTransaction(sgn_tx)
	if err != nil {
		t.Error(err)
	}

	if !ver {
		t.Errorf("Verification failed")
	}
}
