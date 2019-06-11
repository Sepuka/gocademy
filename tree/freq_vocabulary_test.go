package tree

import (
	"strings"
	"testing"

	"github.com/sepuka/gocademy/resources"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	expectedDictionary = `
brother		2
father		1	2	5
mother		1
sister		3
uncle		2	4`
)

type printer struct {
	paper []byte
}

func (pr *printer) Write(p []byte) (n int, err error) {
	pr.paper = append(pr.paper, p...)

	return len(p), nil
}

type lexicographicalOrderBuilder struct {
	suite.Suite
}

func TestLexicographicalOrder(t *testing.T) {
	suite.Run(t, new(lexicographicalOrderBuilder))
}

func (l *lexicographicalOrderBuilder) TestBuild() {
	var (
		printer = &printer{}
		dict *FrequencyVocabularyNode
	)

	dict = BuildFreqVocabulary(resources.TestText)

	Print(dict, printer)
	assert.Equal(
		l.Suite.T(),
		strings.TrimSpace(expectedDictionary),
		strings.TrimSpace(string(printer.paper)),
	)
}
