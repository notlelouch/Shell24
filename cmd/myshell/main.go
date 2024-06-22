package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// trimmed_cons := strings.TrimRight(os.Args[0], "\r\n")
	// trimmed_cons := strings.TrimSuffix(os.Args[0], "\r\n$")

	paths := []string{}
	// path := os.Args[0]

	path := os.Getenv("PATH")
	// /fmt.Println("Value of 'PATH' environment variable:", path)
	pathParts := strings.Split(path, ":")
	paths = append(paths, pathParts...)

	//fmt.Fprint(os.Stdout, paths)

	shellBuiltinCommands := map[string]int{"type": 1, "exit": 1, "pwd": 1, "echo": 1, "cd": 1}
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading", err)
		}
		input = input[:len(input)-1]

		switch {
		case input == "exit 0":
			os.Exit(0)
		case input[:4] == "echo":
			fmt.Fprintf(os.Stdout, "%s\n", input[5:])
		case input[:4] == "type":
			if shellBuiltinCommands[input[5:]] == 1 {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", input[5:])
			} else if shellBuiltinCommands[input[5:]] != 1 {
				for _, path := range paths {
					if _, err := os.Stat(path + "/" + input[5:]); !os.IsNotExist(err) {
						fmt.Fprintf(os.Stdout, "%s is %s\n", input[5:], path+"/"+input[5:])
						break
					}
				}
			} else {
				fmt.Fprintf(os.Stdout, "chuuuuuuuuuuttt")
				fmt.Fprintf(os.Stdout, "%s: not found\n", input[5:])
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
		}

	}
}
