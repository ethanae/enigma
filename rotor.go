package main

import (
	"fmt"
)

type Rotor struct {
	in string
	out string
}

func (r *Rotor) EncodeRL(index int) int {
	outChar := rune(r.out[index])
	for i, v := range r.in {
		if v == outChar {
			fmt.Printf("Enter at %s, exit at %s\n", string(r.in[index]), string(outChar))
			return i
		}
	}

	return -1
}

func (r *Rotor) EncodeLR(index int) int {
	inChar := rune(r.in[index])
	for i, v := range r.out {
		if v == inChar {
			fmt.Printf("Enter at %s, exit at %s\n", string(inChar), string(r.in[i]))
			return i
		}
	}

	return -1
}

func (r *Rotor) Rotate() {
	r.in = r.in[1:] + string(r.in[0])
	r.out = r.out[1:] + string(r.out[0])
}