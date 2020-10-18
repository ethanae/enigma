package main

type Enigma struct {
	rotors    []Rotor
	alphabet  string
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

		didRotateNextRotors := false
		// println("Rotors", string(rotors[0].in[0]), string(rotors[1].in[0]), string(rotors[2].in[0]))
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

		cipher += string(e.alphabet[index])
	}

	return cipher
}

func (e *Enigma) HandleNotchRotations(rotorToTurn int) {
	if rotorToTurn != -1 {
		if e.rotors[rotorToTurn].Rotate() {
			e.HandleNotchRotations(rotorToTurn - 1)
		}
	}
}

func (e *Enigma) Reflect(index int) int {
	reflected := rune(e.reflector[index])
	for i, v := range e.alphabet {
		if reflected == v {
			return i
		}
	}
	return -1
}
