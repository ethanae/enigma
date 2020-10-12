package main

import "testing"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestSetRingByLetterA(t *testing.T) {
	const answer = "FLNGMHERWAOUPXZIYVTQBJCSDK"
	const ringSetting = "B"
	rotor := Rotor{ wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ" }

	rotor.SetRingByLetter(ringSetting, alphabet)

	if rotor.ringSetting != ringSetting {
		t.Errorf("Got %s, expected %s", rotor.ringSetting, ringSetting)
	} else if rotor.wiring != answer {
		t.Errorf("Got %s, expected %s", rotor.wiring, answer)
	}
}

func TestSetRingByLetterC(t *testing.T) {
	const answer = "GMOHNIFSXBPVQYAJZWURCKDTEL"
	const ringSetting = "C"
	rotor := Rotor{ wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ" }

	rotor.SetRingByLetter(ringSetting, alphabet)
	
	if rotor.ringSetting != ringSetting {
		t.Errorf("Got %s, expected %s", rotor.ringSetting, ringSetting)
	} else if rotor.wiring != answer {
		t.Errorf("Got %s, expected %s", rotor.wiring, answer)
	}
}

func TestEncryptA(t *testing.T) {
	expected := "E"
	rotor := Rotor{ wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ" }

	actual := rotor.Encrypt("A", alphabet)

	if actual != expected {
		t.Errorf("Expected %s, but received %s", expected, actual)
	}
}

func TestEncryptZ(t *testing.T) {
	expected := "J"
	rotor := Rotor{ wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ" }

	actual := rotor.Encrypt("Z", alphabet)

	if actual != expected {
		t.Errorf("Expected %s, but received %s", expected, actual)
	}
}