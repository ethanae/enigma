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
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
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
		alphabet:  alphabet,
		reflector: reflector,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
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
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
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
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
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
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
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
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
	}

	actual := enigma.EncryptMessage(message)

	expected := "HELLO"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineEncryptTripleNotchRotation(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"Q",
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
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
	}

	actual := enigma.EncryptMessage(message)

	expected := "FFETW"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineDecryptTripleNotchRotation(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"Q",
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

	message := "FFETW"
	enigma := Enigma{
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
	}

	actual := enigma.EncryptMessage(message)

	expected := "HELLO"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineEncryptLongText(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"Q",
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

	message := "THEQUICKBROWNFOXJUMPSOVERTHELAZYDOGTHEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"
	enigma := Enigma{
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
	}

	actual := enigma.EncryptMessage(message)

	expected := "KVLJBUFICCBYIOIODPWHQTQJZVSTFZXEEYSNODWVQBWAHSIUZFYQFIYFKADTIKTPNLURQK"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineDecryptLongText(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"Q",
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

	message := "KVLJBUFICCBYIOIODPWHQTQJZVSTFZXEEYSNODWVQBWAHSIUZFYQFIYFKADTIKTPNLURQK"
	enigma := Enigma{
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
	}

	actual := enigma.EncryptMessage(message)

	expected := "THEQUICKBROWNFOXJUMPSOVERTHELAZYDOGTHEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineEncryptExtraLongText(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"X",
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"G",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"Z",
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"W",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"N",
		"BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"Z",
		"W",
	)

	reflectorB := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	message := "SEDUTPERSPICIATISUNDEOMNISISTENATUSERRORSITVOLUPTATEMACCUSANTIUMDOLOREMQUELAUDANTIUMTOTAMREMAPERIAMEAQUEIPSAQUAEABILLOINVENTOREVERITATISETQUASIARCHITECTOBEATAEVITAEDICTASUNTEXPLICABONEMOENIMIPSAMVOLUPTATEMQUIAVOLUPTASSITASPERNATURAUTODITAUTFUGITSEDQUIACONSEQUUNTURMAGNIDOLORESEOSQUIRATIONEVOLUPTATEMSEQUINESCIUNTNEQUEPORROQUISQUAMESTQUIDOLOREMIPSUMQUIADOLORSITAMETCONSECTETURADIPISCIVELITSEDQUIANONNUMQUAMEIUSMODITEMPORAINCIDUNTUTLABOREETDOLOREMAGNAMALIQUAMQUAERATVOLUPTATEMUTENIMADMINIMAVENIAMQUISNOSTRUMEXERCITATIONEMULLAMCORPORISSUSCIPITLABORIOSAMNISIUTALIQUIDEXEACOMMODICONSEQUATURQUISAUTEMVELEUMIUREREPREHENDERITQUIINEAVOLUPTATEVELITESSEQUAMNIHILMOLESTIAECONSEQUATURVELILLUMQUIDOLOREMEUMFUGIATQUOVOLUPTASNULLAPARIATUR"
	enigma := Enigma{
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
	}

	actual := enigma.EncryptMessage(message)
	expected := "LTQLROOOINCLFPZAKZTVDGBPGXPAQMJUMGZNKVDJHLGHGCXEETPNUMOLNZHFMXAQSTPGNJXVKPQUFZJJJCAPDJQIEQWSIRRNMZUCVSAJYLRLSBHKQDQGUCEUZBIMKKUSCUGXMPUAQOUKQNJIWKWAVUOQHMZKXFFTSMCJOTLDEPDUMCBMGJALCYQYWEXRSOBWBUXAFCHURYRRWHHQLLCXNMKVQHTQCEIVDOWWOBDXZVHYSHRIPJIKQYGKSXTKLVZGWMROWNDLZZXGYSANTZJWRPXBLHQRNJPRYKJPTRKGPORNNLWCEGJADPLQDXCQAGUZFUBJSZGMSFTPPGIPBHUDBSENYEERUBZTBFCTPOOEOKGVNQRRTBHOVSVXECZBVBYRJDBLOTSEIJGGKXISVUJZPBCEPRDVJOOXUVLFHLOFKVXLCCYSMIPURQLQQRZTAHQCCVROODZBLCSSDKTWJQKVTVFSBGVFUFWSVKJVVOXUSCYCQBJTCBFHREWTNYUMKRWJVUBWWJLTBZUDJNYSLFEEVBBKGBYEMMMCUWHTNVTCYWVWEAOCCYWSIPKNQXDAXEHMETYDVBYNOOCTCRCJYACCKBWYRNTYFQTZLTRLLMXXGKZZUYRVYZDXHWTATJMOGQGUQKGSEPRRFYNTATDFZFVXDWMHQMOMVOLWOMOGGZHYPKUDYTIYJIHRQMBTYRMZEYCTVNIHLLMCCEIIIZYMLM"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineDecryptExtraLongText(t *testing.T) {
	leftRotor := NewRotor(
		alphabet,
		"X",
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"G",
		"R",
	)
	middleRotor := NewRotor(
		alphabet,
		"Z",
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"W",
		"F",
	)
	rightRotor := NewRotor(
		alphabet,
		"N",
		"BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"Z",
		"W",
	)

	reflectorB := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	message := "LTQLROOOINCLFPZAKZTVDGBPGXPAQMJUMGZNKVDJHLGHGCXEETPNUMOLNZHFMXAQSTPGNJXVKPQUFZJJJCAPDJQIEQWSIRRNMZUCVSAJYLRLSBHKQDQGUCEUZBIMKKUSCUGXMPUAQOUKQNJIWKWAVUOQHMZKXFFTSMCJOTLDEPDUMCBMGJALCYQYWEXRSOBWBUXAFCHURYRRWHHQLLCXNMKVQHTQCEIVDOWWOBDXZVHYSHRIPJIKQYGKSXTKLVZGWMROWNDLZZXGYSANTZJWRPXBLHQRNJPRYKJPTRKGPORNNLWCEGJADPLQDXCQAGUZFUBJSZGMSFTPPGIPBHUDBSENYEERUBZTBFCTPOOEOKGVNQRRTBHOVSVXECZBVBYRJDBLOTSEIJGGKXISVUJZPBCEPRDVJOOXUVLFHLOFKVXLCCYSMIPURQLQQRZTAHQCCVROODZBLCSSDKTWJQKVTVFSBGVFUFWSVKJVVOXUSCYCQBJTCBFHREWTNYUMKRWJVUBWWJLTBZUDJNYSLFEEVBBKGBYEMMMCUWHTNVTCYWVWEAOCCYWSIPKNQXDAXEHMETYDVBYNOOCTCRCJYACCKBWYRNTYFQTZLTRLLMXXGKZZUYRVYZDXHWTATJMOGQGUQKGSEPRRFYNTATDFZFVXDWMHQMOMVOLWOMOGGZHYPKUDYTIYJIHRQMBTYRMZEYCTVNIHLLMCCEIIIZYMLM"
	enigma := Enigma{
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
	}

	actual := enigma.EncryptMessage(message)
	expected := "SEDUTPERSPICIATISUNDEOMNISISTENATUSERRORSITVOLUPTATEMACCUSANTIUMDOLOREMQUELAUDANTIUMTOTAMREMAPERIAMEAQUEIPSAQUAEABILLOINVENTOREVERITATISETQUASIARCHITECTOBEATAEVITAEDICTASUNTEXPLICABONEMOENIMIPSAMVOLUPTATEMQUIAVOLUPTASSITASPERNATURAUTODITAUTFUGITSEDQUIACONSEQUUNTURMAGNIDOLORESEOSQUIRATIONEVOLUPTATEMSEQUINESCIUNTNEQUEPORROQUISQUAMESTQUIDOLOREMIPSUMQUIADOLORSITAMETCONSECTETURADIPISCIVELITSEDQUIANONNUMQUAMEIUSMODITEMPORAINCIDUNTUTLABOREETDOLOREMAGNAMALIQUAMQUAERATVOLUPTATEMUTENIMADMINIMAVENIAMQUISNOSTRUMEXERCITATIONEMULLAMCORPORISSUSCIPITLABORIOSAMNISIUTALIQUIDEXEACOMMODICONSEQUATURQUISAUTEMVELEUMIUREREPREHENDERITQUIINEAVOLUPTATEVELITESSEQUAMNIHILMOLESTIAECONSEQUATURVELILLUMQUIDOLOREMEUMFUGIATQUOVOLUPTASNULLAPARIATUR"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}

func TestMachineEncryptSimpleWithPlugboard(t *testing.T) {
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
		alphabet:  alphabet,
		reflector: reflectorB,
		rotors:    []Rotor{leftRotor, middleRotor, rightRotor},
		plugboard: Plugboard{
			in:  "BCDEKMOPUG",
			out: "QRIJWTSXZH",
		},
	}

	actual := enigma.EncryptMessage(message)

	expected := "CAUSC"
	if actual != expected {
		t.Errorf("Expected %s, received %s", string(expected), string(actual))
	}
}
