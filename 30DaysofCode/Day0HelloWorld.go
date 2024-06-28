package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// declaring the variable using the var keyword
	var inputString string

	// scanning the input by the user
	inputReader := bufio.NewReader(os.Stdin)
	inputString, _ = inputReader.ReadString('\n')

	fmt.Print("Hello, World.\n", inputString)
}
