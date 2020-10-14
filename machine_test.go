package main

import (
	"testing"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestMachineEncrypt1(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"A",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"A",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"A",
		"W",
	)

	reflectorB := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	message := "AAAAA"
	enigma := Enigma{
		alphabet: alphabet,
		reflector: reflectorB,
		rotors: []Rotor{ leftRotor, middleRotor, rightRotor, },
	}
	
	actual := enigma.EncryptMessage(message)

	expected := "BDZGO"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineEncrypt2(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"A",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"Y",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"I",
		"W",
	)

	reflector := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	message := "QWERT"
	enigma := Enigma{
		alphabet: alphabet,
		reflector: reflector,
		rotors: []Rotor{ leftRotor, middleRotor, rightRotor, },
	}
	
	actual := enigma.EncryptMessage(message)

	expected := "JVKTO"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}
