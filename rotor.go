package main

import (
	"fmt"
)

type Rotor struct {
	in string
	out string
	turnover string
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

func (r *Rotor) Rotate() bool {
	firstAlphaChar := string(r.in[0])
	r.in = r.in[1:] + firstAlphaChar
	r.out = r.out[1:] + string(r.out[0])

	return firstAlphaChar == r.turnover
}

func NewRotor(alphabet string, wiring string, ringSetting string, turnover string) Rotor {
	var shift int32 = 0
	for _, value := range alphabet {
		if string(value) == ringSetting {
			break
		}
		shift++
	}
	shiftedWiring := ""
	alphabetLen := len(alphabet)
	firstAlphabetRune := rune(alphabet[0])
	dotPosition := 0

	for i, value := range wiring {
		if string(value) == "A" {
			dotPosition = (i + int(shift)) % alphabetLen
		}
		index := (value - firstAlphabetRune + shift) % int32(alphabetLen)
		println("index", index, "char", string(value))
		shiftedChar := string(alphabet[index])
		shiftedWiring += shiftedChar
	}

	for string(shiftedWiring[dotPosition]) != ringSetting {
		shiftedWiring = shiftedWiring[1:] + shiftedWiring[0:1]
	}

	return Rotor{
		in: alphabet,
		out: shiftedWiring,
		turnover: turnover,
	}
}