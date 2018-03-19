package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

//REPL implementation
func main() {
	fmt.Println("Welcome to mySQLITE v1 ")

	reader := bufio.NewReader(os.Stdin)

	for {
		//TODO : build event loop here to support querying
		fmt.Print("sqlit> ")
		query, _ := reader.ReadString('\n')

		if strings.Compare(query, "EXIT\n") == 0 {
			fmt.Println("Exiting sqlit!")
			os.Exit(0)
		}

		fmt.Print("Input query : ", query)

	}
}
