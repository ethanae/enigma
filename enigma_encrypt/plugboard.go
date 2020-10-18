package main

// Plugboard mimics the structure of a the plugboard of an Enigma machine
// It bi-directionally maps two letters to one another
type Plugboard struct {
	in  string
	out string
}

// Map maps an input letter to an output letter based on the Plugboard configuration
func (pb *Plugboard) Map(letter string) string {
	for i, l := range pb.in {
		if string(l) == letter {
			return string(pb.out[i])
		}
	}
	for i, l := range pb.out {
		if string(l) == letter {
			return string(pb.in[i])
		}
	}
	return letter
}
