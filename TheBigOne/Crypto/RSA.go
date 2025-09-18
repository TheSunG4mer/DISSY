package RSAandAES

import (
	"crypto/rand"
	"math/big"
)

func KeyGen(k int) (*big.Int, *big.Int, *big.Int, error) {
	r := big.NewInt(1)
	one := big.NewInt(1)
	two := big.NewInt(2)
	three := big.NewInt(3)

	var p, q *big.Int
	var err error

	for {
		p, err = rand.Prime(rand.Reader, k)
		if err != nil {
			return nil, nil, nil, err
		}
		r.Mod(p, three)
		if r.Cmp(two) == 0 {
			break
		}
	}

	for {
		q, err = rand.Prime(rand.Reader, k)
		if err != nil {
			return nil, nil, nil, err
		}
		r.Mod(q, three)
		if r.Cmp(two) == 0 {
			break
		}
	}

	// fmt.Println(p, q)
	n := big.NewInt(1)
	n.Mul(p, q)

	phi_n := big.NewInt(1)
	phi_n.Mul(p.Sub(p, one), q.Sub(q, one))
	// fmt.Println("KeyGen")

	d := big.NewInt(1)
	d.ModInverse(three, phi_n)

	return n, three, d, nil
}

func Encrypt(msg *big.Int, n *big.Int, e *big.Int) *big.Int {
	c := big.NewInt(1)
	c.Exp(msg, e, n)
	return c
}

func Decrypt(cif *big.Int, n *big.Int, d *big.Int) *big.Int {
	p := big.NewInt(1)
	p.Exp(cif, d, n)
	return p
}
