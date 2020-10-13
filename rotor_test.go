package main

import (
	"testing"
)

func TestRotateAndEncodeLRA(t *testing.T) {
	r := Rotor{
		in: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		out: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	}

	outIndex := r.EncodeLR(3)

	expected := alphabet[1]
	actual := alphabet[outIndex]
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestRotateTwiceAndEncodeLRA(t *testing.T) {
	r := Rotor{
		in: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		out: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	}

	r.Rotate()
	r.Rotate()
	outIndex := r.EncodeLR(25)

	expected := alphabet[24]
	actual := alphabet[outIndex]
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestRotateTwiceAndEncodeRLA(t *testing.T) {
	r := Rotor{
		in: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		out: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
	}

	r.Rotate()
	outIndex1 := r.EncodeRL(0)
	r.Rotate()
	outIndex2 := r.EncodeRL(0)
	
	expected1 := alphabet[2]
	actual1 := alphabet[outIndex1]
	if actual1 != expected1 {
		t.Errorf("Expected %s, received %s", string(expected1), string(actual1))
	}

	expected2 := alphabet[3]
	actual2 := alphabet[outIndex2]
	if actual2 != expected2 {
		t.Errorf("Expected %s, received %s", string(expected2), string(actual2))
	}
}