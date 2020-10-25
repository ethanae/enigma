```
███████╗███╗   ██╗██╗ ██████╗ ███╗   ███╗ █████╗ 
██╔════╝████╗  ██║██║██╔════╝ ████╗ ████║██╔══██╗
█████╗  ██╔██╗ ██║██║██║  ███╗██╔████╔██║███████║
██╔══╝  ██║╚██╗██║██║██║   ██║██║╚██╔╝██║██╔══██║
███████╗██║ ╚████║██║╚██████╔╝██║ ╚═╝ ██║██║  ██║
```
### A badly written Enigma machine, but it works 🤷‍♂️

### Usage
You'll need a file which describes the configuration or key for the machine.

#### Configuration
Every configuration file must have:
- a rotors array with all fields required for a rotor
- an alphabet
- a reflector
- plugboard (optional)

An example configuration file:
```
{
  "alphabet": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
  "reflector": "YRUHQSLDPXNGOKMIEBFZCWVJAT", // reflector B
  "plugboard": "AB, FG, PZ", // A -> B and B -> A, P -> Z and Z -> Z, etc. Either an empty string, or a list of unique non-overlapping pairs
  "rotors": [
    {
      "start": "Q",
      "wiring": "EKMFLGDQVZNTOWYHXUSPAIBRCJ", // Rotor 1
      "ring": "E",
      "turnoverLetter": "R"
    },
    {
      "start": "E",
      "wiring": "AJDKSIRUXBLHWTMCQGZNPYFVOE", // Rotor 2
      "ring": "X",
      "turnoverLetter": "F"
    },
    {
      "start": "V",
      "wiring": "BDFHJLCPRTXVZNYEIWGAKMUSQO", // Rotor 3
      "ring": "K",
      "turnoverLetter": "W"
    }
  ]
}
```

#### Encrypt/descrypt things
Run the using python 3:

```
python main.py --config path/to/configuration/file.json 
```
