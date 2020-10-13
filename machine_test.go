package main

import (
	"testing"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestMachineEncryptAAAAA(t *testing.T) {
	r1 := Rotor{
		in: alphabet,
		out: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	}
	r2 := Rotor{
		in: alphabet,
		out: "AJDKSIRUXBLHWTMCQGZNPYFVOE",
	}
	r3 := Rotor{
		in: alphabet,
		out: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	}

	reflector := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	message := "AAAAA"
	enigma := Enigma{
		alphabet: alphabet,
		reflector: reflector,
		rotors: []Rotor{ r1, r2, r3, },
	}
	
	actual := enigma.EncryptMessage(message)

	expected := "BDZGO"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}
