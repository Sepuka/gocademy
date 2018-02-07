package main

import (
	"os"
	"fmt"
	"strings"
	"gocademy/search"
)

const (
	Linear = "linear"
)

var AvailableFuncs = []string{Linear}

func main() {
	var funcName = getFuncName()
	funcArgsValidate(funcName)
	switch funcName {
		default:
			fmt.Printf("Bad func name %v\n", funcName)
			fmt.Printf("usage: %v %v", os.Args[1], avalableFuncs())
			os.Exit(1)
		case Linear:
			var pattern = os.Args[2]
			var symbol = os.Args[3]
			var result, err = linear.LinearSearch(pattern, symbol)
			if err == nil {
				fmt.Printf("Symbol was found with offset %v!", result)
			} else {
				fmt.Printf("Search error for symbol '%v': %v", symbol, err)
			}
	}
}

func getFuncName() (funcName string) {
	if len(os.Args) < 2 {
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

func avalableFuncs() (string) {
	return strings.Join(AvailableFuncs, " | ")
}