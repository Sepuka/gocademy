package main

import (
	"os"
	"fmt"
	"errors"
)

func main() {
	if getFuncName() == "linear" {
		var result, err = LinearSearch(os.Args[2], os.Args[3])
		if err == nil {
			fmt.Printf("Symbol was found with offset %v!", result)
		} else {
			fmt.Printf("Linear error: %v", err)
		}
	} else {
		fmt.Println("Bad func name %v", os.Args[0])
		os.Exit(1)
	}
}

func getFuncName() (funcName string) {
	if len(os.Args) == 0 {
		fmt.Printf("Expected func name")
		os.Exit(1)
	}

	return os.Args[1]
}

func LinearSearch(data string, symbol string) (result int, err error) {
	for current := range data {
		if symbol == string(data[current]) {
			return current, nil
		}
	}

	return -1, errors.New("Symbol not found")
}