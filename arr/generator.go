package arr

import (
	"errors"
	"sync"
)

// Generator is a thread-safe generator of integer numbers.
type Generator interface {
	Next() (int, bool)
	Err() error
}

// genStruct implements interface Generator.
type genStruct struct {
	sync.Mutex
	Start int
	Stop  int
	Step  int
	cur   int
	err   error
}

func (g *genStruct) Next() (int, bool) {
	if g.cur >= g.Stop {
		return 0, false
	}
	if g.err != nil {
		return 0, false
	}

	g.Lock()
	var r = g.cur
	g.cur += g.Step
	g.Unlock()

	return r, true
}

// Err returns an error if it occurred during generation.
func (g *genStruct) Err() error {
	if g.err != nil {
		return g.err
	}

	return nil
}

// New returns new Generator interface.
func New(start, stop, step int) Generator {
	var err error
	if step <= 0 {
		err = errors.New("failed step")
	}
	if start > stop {
		err = errors.New("failed start/stop")
	}

	return &genStruct{Start: start, Stop: stop, Step: step, cur: start, err: err}
}
