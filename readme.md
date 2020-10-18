```
███████╗███╗   ██╗██╗ ██████╗ ███╗   ███╗ █████╗ 
██╔════╝████╗  ██║██║██╔════╝ ████╗ ████║██╔══██╗
█████╗  ██╔██╗ ██║██║██║  ███╗██╔████╔██║███████║
██╔══╝  ██║╚██╗██║██║██║   ██║██║╚██╔╝██║██╔══██║
███████╗██║ ╚████║██║╚██████╔╝██║ ╚═╝ ██║██║  ██║
```
### A badly written Enigma machine

### Build
You'll need the Golang build tools to build an executable.

#### Linux
In the root of the project run:
```
env GOOS=linux GOARCH=amd64 go build -o ../build/enigma.sh
```

#### Windows
In the root of the project run:
```
env GOOS=windows GOARCH=amd64 go build -o ../build/enigma.exe 
```

### Usage
You also need a file which describes the configuration or key for the machine.

#### Configuration
An example configuration file:
```
{
  "alphabet": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
  "reflector": "YRUHQSLDPXNGOKMIEBFZCWVJAT", // reflector B
  "plugboard": "AB, FG, PZ", // A -> B and B -> A, P -> Z and Z -> Z, etc. Either empty, or a list of unique non-overlapping pairs
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

#### Encrypt things
Run the executable:

```
../build/enigma.sh -c path/to/configuration/file.json 
```
Or, for Windows:

```
../build/enigma.exe -c path/to/configuration/file.json 
```
