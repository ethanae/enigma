package main

import (
	"fmt"
)

type Enigma struct {
	rotors []Rotor
	alphabet string
	reflector string
}

func (e *Enigma) EncryptMessage(message string) string {
	
	rotors := e.rotors
	rotorCount := len(rotors)
	cipher := ""

	for _, letter := range message {
		index := 0
		for j, c := range e.alphabet {
			if string(c) == string(letter) {
				index = j
				break
			}
		}

		for i := rotorCount - 1; i > -1; i-- {
			if i == rotorCount - 1 {
				rotors[i].Rotate()
			}
			index = rotors[i].EncodeRL(index)
		}

		index = e.Reflect(index)

		for _, rotorX := range e.rotors  {
			index = rotorX.EncodeLR(index)
		}

		cipher += string(e.alphabet[index])
	}

	return cipher
}

func (e *Enigma) Reflect(index int) int {
	reflected := rune(e.reflector[index])
	for i, v := range e.alphabet {
		if reflected == v {
			fmt.Printf("Enter at %s, reflected as %s \n", string(reflected), string(e.alphabet[i]))
			return i
		}
	}
	return -1;
}
