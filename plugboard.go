package main

type Plugboard struct {
	in  string
	out string
}

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
