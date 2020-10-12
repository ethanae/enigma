package main

import (
	"strings"
)

type Rotor struct {
	ringSetting	string
	notch string
	position string
	wiring string
}

func (rotor *Rotor) Rotate(alphabet string) bool {
	posRune := []rune(rotor.position)[0]
	index := (posRune - 'A' + 1) % int32(len(alphabet))
	rotor.position = string(alphabet[index])
	if rotor.position == rotor.notch {
		return true
	}
	return false
}

func (rotor *Rotor) Encrypt(letter string, alphabet string) string {
	for i, value := range alphabet {
		if letter == string(value) {
			return string(rotor.wiring[i])
		}
	}
	return ""
}

func (rotor *Rotor) SetRingByLetter(letter string, alphabet string) {
	// get this distance of the desired setting provided by "letter"
	// then get the position of "letter" in "alphabet" which is the shift value
	// shift the rotor's wiring based on the shift and modulo to prevent overflows
	rotor.ringSetting = letter
	var shift int32 = 0
	for _, value := range alphabet {
		if string(value) == letter {
			break
		}
		shift++
	}
	shiftedAlphabet := []string{}
	alphabetLen := len(alphabet)
	firstAlphabetRune := rune(alphabet[0])
	for _, value := range rotor.wiring {
		index := (value - firstAlphabetRune + shift) % int32(alphabetLen)
		shiftedChar := string(alphabet[index])
		shiftedAlphabet = append(shiftedAlphabet, shiftedChar)
	}
	rotor.wiring = strings.Join(shiftedAlphabet, "")
}