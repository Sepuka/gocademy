package search

import (
	"fmt"
	"os"
)

const (
	Linear         = "linear"
	ImprovedLinear = "improved_linear"
	SimplePattern  = "simple_pattern"
)

func main() {
	availableFuncsMap := map[string]int{
		Linear:         4,
		ImprovedLinear: 4,
		SimplePattern:  4,
	}

	funcName := getFuncName()
	expectedArgs, funcExists := availableFuncsMap[funcName]
	if !funcExists {
		fmt.Printf("%v func name not exists!\n", funcName)
		os.Exit(1)
	}
	if len(os.Args) != expectedArgs {
		fmt.Printf("%v search usage: %v text symbol\n", funcName, funcName)
		os.Exit(1)
	}

	switch funcName {
	case Linear, ImprovedLinear:
		var pattern = os.Args[2]
		var symbol = os.Args[3]
		var result, err = LinearSearch(pattern, symbol)
		if err == nil {
			fmt.Printf("Symbol was found with offset %v!", result)
		} else {
			fmt.Printf("Search error for symbol '%v': %v", symbol, err)
		}
	case SimplePattern:
		var text = os.Args[2]
		var pattern = os.Args[3]
		var result, err = SimplePatternSearch(text, pattern)
		if err == nil {
			fmt.Printf("Pattern was found with offset %v!", result)
		} else {
			fmt.Printf("Pattern error search for '%v': %v", pattern, err)
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
