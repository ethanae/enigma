package main

import "strings"

type RotorConfiguration struct {
	Start          string `json:"start"`
	Wiring         string `json:"wiring"`
	Ring           string `json:"ring"`
	TurnoverLetter string `json:"turnoverLetter"`
}

type EnigmaConfiguration struct {
	Rotors    []RotorConfiguration `json:"rotors"`
	Alphabet  string               `json:"alphabet"`
	Reflector string               `json:"reflector"`
	Plugboard string               `json:"plugboard"`
}

func NewEnigma(config EnigmaConfiguration) Enigma {
	plugboardPairs := strings.Split(strings.Replace(config.Plugboard, " ", "", -1), ",")
	plugboardIn := ""
	plugboardOut := ""

	for _, pair := range plugboardPairs {
		plugboardIn += string(pair[0])
		plugboardOut += string(pair[1])
	}

	rotors := []Rotor{}

	for _, r := range config.Rotors {
		rotors = append(rotors, NewRotor(
			config.Alphabet,
			r.Start,
			r.Wiring,
			r.Ring,
			r.TurnoverLetter,
		))
	}

	enigma := Enigma{
		alphabet: config.Alphabet,
		plugboard: Plugboard{
			in:  plugboardIn,
			out: plugboardOut,
		},
		reflector: config.Reflector,
		rotors:    rotors,
	}

	return enigma
}
