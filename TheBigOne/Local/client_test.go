package Local_test

import (
	"Local"
	"RSAandAES"
	"testing"
)

func TestStringToPublicKey(t *testing.T) {
	for i := 0; i < 10; i++ {
		n, e, _, err := RSAandAES.KeyGen(200)
		if err != nil {
			t.Error(err)
		}
		pk := Local.PublicKey{n, e}
		// sk := Local.SecretKey{n, d}
		s := Local.PublicKeyToString(&pk)
		pk_ := Local.StringToPublicKey(s)
		if pk.N.Cmp(pk_.N) != 0 {
			t.Errorf("Wrong public key")
		}
		if pk.E.Cmp(pk_.E) != 0 {
			t.Errorf("Wrong public key")
		}
	}
}
