package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("Scanner: " + scanner.Text() + "\n")
	}
}
