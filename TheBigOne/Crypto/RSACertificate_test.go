package RSAandAES

import (
	"testing"
)

func TestCertificate(t *testing.T) {
	n, e, d, err := KeyGen(2000)
	if err != nil {
		t.Errorf("Error occured during key generation")
	}

	msg := []byte("I hope no one copies this message")
	certificate := GenerateCertificate(msg, n, d)

	if !VerifyCertificate(msg, n, e, certificate) {
		t.Errorf("Wrong certificate was given")
	}
}

func TestWrongCertificate(t *testing.T) {
	n, e, d, err := KeyGen(2000)
	if err != nil {
		t.Errorf("Error occured during key generation")
	}

	msg := []byte("I hope no one copies this message")
	certificate := GenerateCertificate(msg, n, d)

	alt_msg := []byte("I hope no one copies this message!")

	if VerifyCertificate(alt_msg, n, e, certificate) {
		t.Errorf("Wrong certificate was given")
	}
}
