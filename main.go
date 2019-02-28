package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Gone: %s", err)
	}
	text = strings.TrimRight(text, "\n")
	tokens := strings.Split(text, " ")
	command := tokens[0]
	if command == "SELECT" {
		fmt.Println(strings.Join(tokens[1:], " "))
	} else {
		log.Fatalf("Unknown command: %s", command)
	}
}
