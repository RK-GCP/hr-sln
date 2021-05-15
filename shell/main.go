package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter a letter(s) seperated by space")
	fmt.Println("enter -1 to end")
	fmt.Println("----------------")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		if input == "-1" {
			os.Exit(0)
		} else {
			if len(input) == 0 {
				input = "-->"
			}
			fmt.Printf("%v\n", input)
		}

	}

}
