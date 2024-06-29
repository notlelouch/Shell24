package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// trimmed_cons := strings.TrimRight(os.Args[0], "\r\n")
	// trimmed_cons := strings.TrimSuffix(os.Args[0], "\r\n$")

	paths := []string{}
	// path := os.Args[0]

	path := os.Getenv("PATH")
	// fmt.Println("Value of 'PATH' environment variable:", path)
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
		//input  = strings.Split(input, " ")

		switch {
		case input == "exit 0":
			os.Exit(0)
		case input == "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintf(os.Stdout, "Error getting current directory")
			} else {
				fmt.Fprintf(os.Stdout, "%s\n", dir)
			}
		case input[:4] == "echo":
			fmt.Fprintf(os.Stdout, "%s\n", input[5:])
		case input[:4] == "type":
			elsebool := false
			if shellBuiltinCommands[input[5:]] == 1 {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", input[5:])
			} else if shellBuiltinCommands[input[5:]] != 1 {
				elsebool = false
				for _, path := range paths {
					if _, err := os.Stat(path + "/" + input[5:]); !os.IsNotExist(err) {
						fmt.Fprintf(os.Stdout, "%s is %s\n", input[5:], path+"/"+input[5:])
						elsebool = true
						break
					}
				}
				if !elsebool {
					fmt.Fprintf(os.Stdout, "%s: not found\n", input[5:])
				}
			}

		case len(strings.Split(input, " ")) > 1:
			program_iput := strings.Split(input, " ")[0]
			argument_input := strings.Split(input, " ")[1]
			for _, path := range paths {
				if _, err := os.Stat(path + "/" + program_iput); !os.IsNotExist(err) {
					cmd := exec.Command(path+"/"+program_iput, argument_input)
					output, err := cmd.CombinedOutput()
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error running %s: %v\n", program_iput, err)
					} else {
						fmt.Fprintf(os.Stdout, "%s", string(output))
					}
					break
				}
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
		}

	}
}
