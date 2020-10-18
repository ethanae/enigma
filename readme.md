## A badly written Enigma machine
#### but it has ASCII art ¯\_(ツ)_/¯

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
  "reflector": "YRUHQSLDPXNGOKMIEBFZCWVJAT",
  "plugboard": "AB, FG, PZ", // A -> B and B -> A, P -> Z and Z -> Z, etc.
  "rotors": [
    {
      "start": "Q",
      "wiring": "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
      "ring": "E",
      "turnoverLetter": "R"
    },
    {
      "start": "E",
      "wiring": "AJDKSIRUXBLHWTMCQGZNPYFVOE",
      "ring": "X",
      "turnoverLetter": "F"
    },
    {
      "start": "V",
      "wiring": "BDFHJLCPRTXVZNYEIWGAKMUSQO",
      "ring": "K",
      "turnoverLetter": "W"
    }
  ]
}
```

#### Encrypt things
With a configuration file at the ready run:

```
../build/enigma.sh -c path/to/configuration/file.json 
```
Or, for Windows:

```
../build/enigma.exe -c path/to/configuration/file.json 
```
