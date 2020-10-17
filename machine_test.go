package main

import (
	"testing"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestMachineEncrypt1(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"E",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"X",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"K",
		"W",
	)

	reflectorB := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	message := "HELLO"
	enigma := Enigma{
		alphabet: alphabet,
		reflector: reflectorB,
		rotors: []Rotor{ leftRotor, middleRotor, rightRotor, },
	}
	
	actual := enigma.EncryptMessage(message)

	expected := "JXZOF"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineDecrypt1(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"E",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"X",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"K",
		"W",
	)

	reflector := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	message := "JXZOF"
	enigma := Enigma{
		alphabet: alphabet,
		reflector: reflector,
		rotors: []Rotor{ leftRotor, middleRotor, rightRotor, },
	}
	
	actual := enigma.EncryptMessage(message)

	expected := "HELLO"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}
