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

//TODO: Create a simple sort function of letters A-Z
//Create a hash table for A - Z
//If list is empty - take it as is
//If list has value - find the index value from has table and insert
// Key Vaule pair may be helpful
