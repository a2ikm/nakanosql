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
	fmt.Println(text)
}
