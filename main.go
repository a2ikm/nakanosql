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
	switch command {
	case "SELECT":
		fmt.Println(strings.Join(tokens[1:], " "))
	case "INSERT":
		f, err := os.OpenFile("data.db", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Can't open data.db: %s", err)
		}
		record := strings.Join(tokens[1:], "\t")
		line := fmt.Sprintln(record)
		if _, err := f.WriteString(line); err != nil {
			log.Fatalf("Can't write data.db: %s", err)
		}
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}
