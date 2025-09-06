package util

import (
	"testing"
	"time"
)

type addTest struct {
	name                 string
	arg1, arg2, expected int
}

var tests = []addTest{
	//{name, arg1, arg2, expected}
	{"basic", 1, 2, 2},
	{"twos", 2, 2, 4},
	{"negative", -1, 2, -2},
	{"negative2", -1, -2, 2},
	{"big", 100, 10, 1000},
}

func TestAdd(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := Mult(test.arg1, test.arg2)
			if res != test.expected {
				t.Errorf("got %d, should be %d", res, test.expected)
			}
		})
	}
}

// run `go test ./util -v -test.short` to skip this test
func TestAddLong(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	time.Sleep(5 * time.Second)
	res := Mult(10, 5)
	if res != 50 {
		t.Errorf("got %d, should be %d", res, 50)
	}
}
