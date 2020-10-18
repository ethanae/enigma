package main

// Rotor models the internals of a physical rotor
// having an alphabet strip, a mapping of wiring, and a turnover notch
type Rotor struct {
	in       string
	out      string
	turnover string
}

// EncodeRL encodes a message based on the right-to-left direction
func (r *Rotor) EncodeRL(index int) int {
	outChar := rune(r.out[index])
	for i, v := range r.in {
		if v == outChar {
			return i
		}
	}

	return -1
}

// EncodeLR encodes a message based on the left-to-right direction
func (r *Rotor) EncodeLR(index int) int {
	inChar := rune(r.in[index])
	for i, v := range r.out {
		if v == inChar {
			return i
		}
	}

	return -1
}

// Rotate simply turns the rotor and returns whether it the next immediate rotor must be turn as well
func (r *Rotor) Rotate() bool {
	r.in = r.in[1:] + string(r.in[0])
	r.out = r.out[1:] + string(r.out[0])
	return string(r.in[0]) == r.turnover
}

// RatchedEngaged is used to determine whether the pawl is engaged by the ratchet on a physical machine
// An engaged ratchet means that the rotor must be turned
func (r *Rotor) RatchedEngaged() bool {
	window := rune(r.turnover[0])
	return string(r.in[0]) == string(window-1)
}

// NewRotor provides the recommended way to create rotors
// Ensures that wiring and starting positions are all adjusted correctly
func NewRotor(
	alphabet string,
	startAt string,
	wiring string,
	ringSetting string,
	turnover string,
) Rotor {
	var shift int32 = 0
	for _, value := range alphabet {
		if string(value) == ringSetting {
			break
		}
		shift++
	}
	shiftedWiring := ""
	alphabetLen := len(alphabet)
	firstAlphabetRune := rune(alphabet[0])
	dotPosition := 0

	for i, value := range wiring {
		if string(value) == "A" {
			dotPosition = (i + int(shift)) % alphabetLen
		}
		index := (value - firstAlphabetRune + shift) % int32(alphabetLen)
		shiftedChar := string(alphabet[index])
		shiftedWiring += shiftedChar
	}

	for string(shiftedWiring[dotPosition]) != ringSetting {
		shiftedWiring = shiftedWiring[1:] + shiftedWiring[0:1]
	}

	rotor := Rotor{
		in:       alphabet,
		out:      shiftedWiring,
		turnover: turnover,
	}

	// shift rotor to the correct starting position
	for string(rotor.in[0]) != startAt {
		rotor.Rotate()
	}

	return rotor
}
