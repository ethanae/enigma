package main

import (
	"errors"
)

// Enigma mimics the internal organisation of a physical Enigma machine
// Each machine has a plugboard, a number of rotors, and a reflector
type Enigma struct {
	rotors    []Rotor
	alphabet  string
	reflector string
	plugboard Plugboard
}

// EncryptMessage takes a full plaintext message and produces the ciphertext determined by the machine's configuration (key)
func (e *Enigma) EncryptMessage(message string) (string, error) {
	rotors := e.rotors
	rotorCount := len(rotors)
	cipher := ""

	for _, letter := range message {
		index := -1
		// pass through plugboard
		nextLetter := e.plugboard.Map(string(letter))

		for j, c := range e.alphabet {
			if string(c) == nextLetter {
				index = j
				break
			}
		}

		if index == -1 {
			return "", errors.New("Invalid input. Character '" + string(letter) + "' not in alphabet '" + e.alphabet + "'")
		}

		didRotateNextRotors := false
		for i := rotorCount - 1; i > -1; i-- {
			if i == rotorCount-1 {
				if rotors[i].Rotate() {
					// turn next rotors
					didRotateNextRotors = true
					e.HandleNotchRotations(i - 1)
				}
			} else if !didRotateNextRotors && rotors[i].RatchedEngaged() {
				if rotors[i].Rotate() {
					// turn next rotors
					e.HandleNotchRotations(i - 1)
				}
			}
			index = rotors[i].EncodeRL(index)
		}

		index = e.Reflect(index)

		for _, rotorX := range e.rotors {
			index = rotorX.EncodeLR(index)
		}

		nextLetter = e.plugboard.Map(string(e.alphabet[index]))

		for i, l := range e.alphabet {
			if string(l) == nextLetter {
				index = i
			}
		}

		cipher += string(e.alphabet[index])
	}

	return cipher, nil
}

// HandleNotchRotations ensures that a rotation of one rotor may end up rotating multiple adjacent rotors
func (e *Enigma) HandleNotchRotations(rotorToTurn int) {
	if rotorToTurn != -1 {
		if e.rotors[rotorToTurn].Rotate() {
			e.HandleNotchRotations(rotorToTurn - 1)
		}
	}
}

// Reflect mimics the behaviour of the reflector on a physical machine which maps an incoming letter to an output letter
func (e *Enigma) Reflect(index int) int {
	reflected := rune(e.reflector[index])
	for i, v := range e.alphabet {
		if reflected == v {
			return i
		}
	}
	return -1
}
