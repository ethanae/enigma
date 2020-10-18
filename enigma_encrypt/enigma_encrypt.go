package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println(`
	███████╗███╗   ██╗██╗ ██████╗ ███╗   ███╗ █████╗ 
	██╔════╝████╗  ██║██║██╔════╝ ████╗ ████║██╔══██╗
	█████╗  ██╔██╗ ██║██║██║  ███╗██╔████╔██║███████║
	██╔══╝  ██║╚██╗██║██║██║   ██║██║╚██╔╝██║██╔══██║
	███████╗██║ ╚████║██║╚██████╔╝██║ ╚═╝ ██║██║  ██║
	╚══════╝╚═╝  ╚═══╝╚═╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝
	`)

	var configFilePath string = ""

	flag.StringVar(
		&configFilePath,
		"c",
		"./config.json",
		"Specify your configuration file. Default is ./config.json",
	)

	flag.Usage = func() {
		fmt.Println("Enigma usage:")
		fmt.Println("enigma.exe -c ./<some_file_path>.json")
	}

	flag.Parse()

	configFile, fileErr := ioutil.ReadFile(configFilePath)
	if fileErr != nil {
		panic("Unable to read file " + configFilePath + ".\n" + fileErr.Error())
	}

	var enigmaConfig EnigmaConfiguration
	err := json.Unmarshal(configFile, &enigmaConfig)
	if err != nil {
		panic("Unable to parse configuration file" + err.Error())
	}

	enigma := NewEnigma(enigmaConfig)
	var plaintext string
	var ciphertext string
	for {
		fmt.Println("------------")
		fmt.Print("plaintext  > ")
		plaintext = readPlainText()
		ciphertext = enigma.EncryptMessage(plaintext)
		fmt.Printf("ciphertext > %s\n", ciphertext)
		fmt.Println("------------")
	}
}

func readPlainText() string {
	text := ""
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()

	return text
}
