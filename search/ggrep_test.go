package search

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const (
	searchPattern = "SeArCh"
)

func writeLine(err error, line string, f *os.File) error {
	if err != nil {
		return err
	}
	line += "\n"
	_, err = f.WriteString(line)
	return err
}

func prepare(pattern string, j int) (string, error) {
	const maxLines = 10
	var line string
	file, err := ioutil.TempFile(os.TempDir(), pattern)
	if err != nil {
		return "", err
	}
	for i := 0; i < maxLines; i++ {
		if i == (j - 1) {
			line = fmt.Sprintf("some text %s other text", searchPattern)
		} else {
			line = fmt.Sprintf("some text, line=%d", i)
		}
		err = writeLine(err, line, file)
	}
	return file.Name(), err
}

func TestSearch(t *testing.T) {
	file1, err := prepare("bcd", 2)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Remove(file1); err != nil {
			t.Error(err)
		}
	}()
	file2, err := prepare("abc", 3)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Remove(file2); err != nil {
			t.Error(err)
		}
	}()
	expected := []Item{
		{Row: 3, FileName: file2},
		{Row: 2, FileName: file1},
	}
	result := Search(searchPattern, []string{file1, file2})
	n, m := len(expected), len(result)
	if n != m {
		t.Fatalf("failed lengths %d != %d", n, m)
	}
	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("failed result [%d] %v != %v", i, expected[i], result[i])
		}
	}
}
