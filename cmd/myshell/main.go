package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var exit bool = false

	shellBuiltinCommands := map[string]int{"type": 1, "exit": 1, "pwd": 1, "echo": 1}
	var pass int = 0
	for {
		if !exit {
			fmt.Fprint(os.Stdout, "$ ")

			input, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading", err)
			}
			input = input[:len(input)-1]

			if input == "exit 0" {
				os.Exit(0)
			}

			if input[:4] == "echo" {
				fmt.Fprintf(os.Stdout, "%s\n", input[5:])
				pass = 1
			}

			if input[:4] == "type" {
				if shellBuiltinCommands[input[5:]] == 1 {
					fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", input[5:])
				} else {
					fmt.Fprintf(os.Stdout, "%s: not found\n", input[5:])
				}
				pass = 1
			}

			if pass != 1 && pass == 0 {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
				pass = 0
			}

		}
	}
}
