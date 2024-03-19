package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nPlease input text below (type '!exit' to quit):")
		text, _ := reader.ReadString('\n')
		trimmedInput := strings.TrimRight(text, "\n")

		if trimmedInput == "!exit" {
			break
		}

		formattedText := strings.Join(strings.Split(strings.ToUpper(trimmedInput), ""), "░")
		fmt.Printf("░%s░\n", formattedText)
	}
}
