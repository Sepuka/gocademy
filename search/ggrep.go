package search

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Item is found result
type Item struct {
	Row      int
	FileName string
}

// ReadFile returns a line number where the pattern was found, otherwise returns 0.
func ReadFile(file, pattern string) (int, error) {
	var n int
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer func() {
		if e := f.Close(); e != nil {
			fmt.Printf("failed close file: %v", err)
		}
	}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n++
		if strings.Contains(line, pattern) {
			return n, nil
		}
	}
	err = scanner.Err()
	if err != nil {
		return 0, err
	}
	return n, nil
}

func Search(pattern string, fileNames []string) []Item {
	var c = make(chan Item, len(fileNames))

	for _, file := range fileNames {
		go func(file string) {
			line, err := ReadFile(file, pattern)
			if err != nil {
				line = 0
			}
			c <- Item{Row: line, FileName: file}
		}(file)
	}

	var res = make([]Item, 0, len(fileNames))

	for range fileNames {
		r := <-c
		if r.Row > 0 {
			res = append(res, r)
		}
	}

	return res
}
