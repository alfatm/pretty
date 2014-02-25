// +build ignore

package main

import (
	"math/big"
	"os"

	"github.com/eaburns/pp"
)

type S struct {
	A, B int
	C    []float64
	D    T
	E    string
	F    *big.Int
	S    *S
}

type T struct {
	X float64
	Y []string
	Z float64
}

func main() {
	s := S{
		A: 5,
		B: 6,
		C: []float64{3.14, 2.8},
		D: T{
			X: 0,
			Y: []string{"foo", "bar"},
			Z: 1.3838,
		},
		E: "Hello, Friends",
		F: big.NewInt(1891284),
	}
	s.S = &s
	err := pp.Print(os.Stdout, &s)
	if err != nil {
		panic(err)
	}
	os.Stdout.WriteString("\n")
}
