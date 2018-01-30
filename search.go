package main

import (
	"os"
	"fmt"
)
const (
	Linear = "linear"
)

func main() {
	var funcName = getFuncName()
	funcArgsValidate(funcName)
	switch funcName {
		default:
			fmt.Printf("Bad func name %v\n", funcName)
			os.Exit(1)
		case Linear:
			var pattern = os.Args[2]
			var symbol = os.Args[3]
			var result = LinearSearch(pattern, symbol)
			if result != -1 {
				fmt.Printf("Symbol was found with offset %v!", result)
			} else {
				fmt.Printf("Symbol '%v' not found", symbol)
			}
	}
}

func getFuncName() (funcName string) {
	if len(os.Args) == 0 {
		fmt.Printf("Expected func name")
		os.Exit(1)
	}

	return os.Args[1]
}

func funcArgsValidate(funcName string) {
	switch funcName {
		default:
			return
		case Linear:
			if len(os.Args) != 4 {
				fmt.Printf("Linear search usage: linear pattern symbol")
				os.Exit(1)
			}
	}
}

func LinearSearch(data string, symbol string) (result int) {
	for current := range data {
		if symbol == string(data[current]) {
			return current
		}
	}

	return -1
}