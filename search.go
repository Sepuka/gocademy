package main

import (
	"os"
	"fmt"
	"gocademy/search"
)

const (
	Linear = "linear"
)

func main() {
	availableFuncss := map[string]int{
		Linear: 4,
	}

	funcName := getFuncName()
	expectedArgs, funcExists := availableFuncss[funcName]
	if !funcExists {
		fmt.Printf("%v func name not exists!\n", funcName)
		os.Exit(1)
	}
	if len(os.Args) != expectedArgs {
		fmt.Printf("%v search usage: %v pattern symbol\n", funcName, funcName)
		os.Exit(1)
	}

	switch funcName {
		case Linear:
			var pattern = os.Args[2]
			var symbol = os.Args[3]
			var result, err = search.LinearSearch(pattern, symbol)
			if err == nil {
				fmt.Printf("Symbol was found with offset %v!", result)
			} else {
				fmt.Printf("Search error for symbol '%v': %v", symbol, err)
			}
	}
}

func getFuncName() (funcName string) {
	if len(os.Args) < 2 {
		fmt.Printf("Expected func name!\n")
		os.Exit(1)
	}

	return os.Args[1]
}