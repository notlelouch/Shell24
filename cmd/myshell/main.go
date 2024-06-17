package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading", err)
	}

	input = input[:len(input)-1]

	fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
}
