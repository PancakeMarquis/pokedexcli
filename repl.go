package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		fmt.Printf("Your command was: %s \n", words[0])

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words

}
