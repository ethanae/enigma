package main

import (
	"testing"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestMachineEncryptSimple(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"A",
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"E",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"A",
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"X",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"A",
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

func TestMachineDecryptSimple(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"A",
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"E",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"A",
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"X",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"A",
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

func TestMachineEcryptSingleNotchRotation(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"T",
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"E",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"L",
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"X",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"V",
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

	expected := "IAWWT"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineDecryptSingleNotchRotation(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"T",
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"E",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"L",
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"X",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"V",
		"BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"K",
		"W",
	)

	reflectorB := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	message := "IAWWT"
	enigma := Enigma{
		alphabet: alphabet,
		reflector: reflectorB,
		rotors: []Rotor{ leftRotor, middleRotor, rightRotor, },
	}
	
	actual := enigma.EncryptMessage(message)

	expected := "HELLO"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineEncryptDoubleNotchRotation(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"A",
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"E",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"E",
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"X",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"V",
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

	expected := "NSAFN"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineDecryptDoubleNotchRotation(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"A",
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"E",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"E",
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"X",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"V",
		"BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"K",
		"W",
	)

	reflectorB := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	message := "NSAFN"
	enigma := Enigma{
		alphabet: alphabet,
		reflector: reflectorB,
		rotors: []Rotor{ leftRotor, middleRotor, rightRotor, },
	}
	
	actual := enigma.EncryptMessage(message)

	expected := "HELLO"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}